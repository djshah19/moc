function Install-AzTools {
    param(
        [parameter(Mandatory = $false, HelpMessage = "Install AZCli.")]
        [switch]$InstallAzCli
    )
    # install AZ Cli if didn't/ 
    if ( $InstallAzCli.IsPresent) {
        $azCliPath = "C:\Program Files (x86)\Microsoft SDKs\Azure\CLI2\wbin"
        if ($false -eq (Test-Path -Path $azCliPath)) {
            Write-Host "Didn't find AZ CLI Core in $azCliPath, download and install AZ CLI ..."
            Invoke-WebRequest -Uri https://aka.ms/installazurecliwindows -OutFile .\AzureCLI.msi 
            Start-Process msiexec.exe -Wait -ArgumentList '/I AzureCLI.msi /quiet'
            Remove-Item .\AzureCLI.msi
            Write-Host "AZ CLI installed."
        }

        if (!(($env:PATH).Split(';') -contains $azCliPath)) {
            $env:PATH += ";$azCliPath;"    
        }
    }

    $azcopyInstallPath = "c:\Azcopy"
    $azcopyPath = Join-Path $azcopyInstallPath "azcopy.exe"

    if ($false -eq (Test-Path -Path $azcopyPath)) {
        Write-Host "Didn't find AZ copy in $azcopyPath, download and install AZCopy ..."
        $zipPath = Join-Path $azcopyInstallPath "azcopy.zip"
        New-Item -Path $zipPath -Force
        #Start-BitsTransfer -Source https://azcopyvnext.azureedge.net/release20211027/azcopy_windows_amd64_10.13.0.zip -Destination $zipPath |Wait-Process
        Start-BitsTransfer -Source https://aka.ms/downloadazcopy-v10-windows -Destination $zipPath | Wait-Process
        Expand-Archive -LiteralPath $zipPath -DestinationPath $azcopyInstallPath
        Remove-Item $zipPath -Force

        $exeFile = Get-ChildItem -Path $azcopyInstallPath -Filter "azcopy.exe" -Recurse
    
        # move the file to root folder.
        Move-Item -Path $exeFile.FullName -Destination $azcopyPath -Force
        Write-Host "AZCopy installed."
    }

    if (!(($env:PATH).Split(';') -contains $azcopyInstallPath)) {
        $env:PATH += ";$azcopyInstallPath"
    }
}

function Invoke-AzCli {
    param(

        [parameter(Mandatory = $true)]
        [string]$script,
        [parameter(Mandatory = $false)]
        [switch]$AsJson
    )

    $ErrorActionPreference = "Continue"

    $jsonOutputArg = if ($AsJson) {
        "--output json"
    }

    $scriptBlock = [scriptblock]::Create("$script $jsonOutputArg 2>&1")

    if ($MyInvocation.ExpectingInput) {
        Write-Verbose "Invoking with input: $script"
        $output = Invoke-Command $scriptBlock 2>&1
    }
    else {
        Write-Verbose "Invoking: $script"
        $output = Invoke-Command $scriptBlock
    }
    
    if ($LASTEXITCODE) {
        # ingore some known error.
        $errorRecords = $output | Where-Object { $_.gettype().Name -eq "ErrorRecord" } | Where-Object { $_.Exception.Message -notmatch "Please let us know how we are doing" `
                -and $_.Exception.Message -notmatch "The installed extension '.+' is experimental" `
                -and $_.Exception.Message -notmatch "The installed extension '.+' is in preview." `
                -and $_.Exception.Message -notmatch "Setting GA feature gate arcmonitoring=true" `
                -and $_.Exception.Message -notmatch "Command group '.+' is in preview and under development." `
                -and $_.Exception.Message -notmatch "ResourceNotFound" `
                -and $_.Exception.Message -notmatch "ARMResourceNotFoundFix" `
                -and $_.Exception.Message -notmatch "could not be found" `
                -and $_.Exception.Message -notmatch "You are using cryptography on a 32-bit Python on a 64-bit Windows Operating System." }
        
        # For show command, return 3 means not found.
        if (($LASTEXITCODE -ne 3) -and ($null -ne $errorRecords)) {
            $stack = Get-PSCallStack
            # An error message was returned, just throw that message
            $errMessage = "$command $arguments returned a non empty error stream [$errorRecords] at [$($stack)]"
            throw $errMessage
        }
        
        # otherwise, all failure are known error, ignore it.
        $out = $output | Where-Object { $_.gettype().Name -ine "ErrorRecord" }  # On a non-zero exit code, this may contain the error
        return $out
        
    }
    else {

        $out = $output | Where-Object { $_.gettype().Name -ine "ErrorRecord" }  # On a non-zero exit code, this may contain the error
        if ($AsJson) {
            return $out | ConvertFrom-Json
        }
        else {
            return $out
        }
    }
}

function Ensure-CleanFolder {
    param (
        [string]$folderPath
    )

    # Clean the $folderPath folder if exists.
    if (Test-Path -Path $folderPath) {
        Remove-Item -Path $folderPath\* -Recurse -Force
    }

    # Create the folder if it doesn't exist
    if (-not (Test-Path -Path $folderPath)) {
        New-Item -Path $folderPath -ItemType Directory
    }
}

Install-AzTools -InstallAzCli

# Set install extension automatically
Invoke-AzCli -script "az config set extension.use_dynamic_install=yes_without_prompt"

$configPath = "C:\CloudDeployment\BVTs\configuration.json"
$config = Get-Content -Path $configPath -Raw | ConvertFrom-Json
$registrationAccountName = $config.AzureRegistrationAccountId
$registrationAccountPassword = $config.AzureRegistrationPassword
$registrationSpAccountName = $config.RegistrationSPAppId
$registrationSPAccountPassword = $config.RegistrationSPPassword
$domain = $config.DomainAdminUser.Split(".")[0]
$hostAccountName = "$domain\$($config.HCIAdministrator)"
$hostAccountPassword = $config.CloudAdminPassword
$subscriptionId = $config.AzureRegistrationSubscriptionId

$spAppId = $config.AKSServicePrincipalID
$spPassword = $config.AKSServicePrincipalPassword
$tenantName = $config.AzureRegistrationDirectoryTenantName

if ($tenantName -eq "aszregistration.onmicrosoft.com") {
    $spTenant = "d9b73d5e-a9d3-41ba-88c3-796a643e3edd"
}
else {
    # Default is microsoft.onmicrosoft.com.
    $spTenant = "72f988bf-86f1-41af-91ab-2d7cd011db47"
}

$deployConfigPath = "C:\CloudDeployment\DeploymentData\DeploymentData.json"
$deployConfig = Get-Content -Path $deployConfigPath -Raw | ConvertFrom-Json

$hostLists = @($deployConfig.ScaleUnits[0].DeploymentData.PhysicalNodes | Select-Object -ExpandProperty Name)

# new-credential is a command in cloudcommon module.
$remoteCred = New-Credential -UserName $hostAccountName -Password $hostAccountPassword

function DoAzLogin {
    # ATTENTION: If you updated this method, don't forget to update the duplicated method in Get-HostPsSession
    #install AZ Cli if didn't/ 
    Install-AzTools -InstallAzCli

    if (($false -eq [string]::IsNullOrEmpty($registrationAccountName)) -and ($false -eq [string]::IsNullOrEmpty($registrationAccountPassword))) {
        if ($false -eq $registrationAccountName.contains("@")) {
            Write-Host "Try to logon with spn"
            $loginResult = az login --service-principal -u $registrationAccountName -p $registrationAccountPassword --tenant $spTenant | ConvertTo-Json
        }
        else {
            Write-Host "Try to logon with username/password"
            $loginResult = az login -u $registrationAccountName -p $registrationAccountPassword --tenant $spTenant | ConvertTo-Json
        }
    }

    if (($null -eq $loginResult) -and ($false -eq [string]::IsNullOrEmpty($registrationSpAccountName)) -and ($false -eq [string]::IsNullOrEmpty($registrationSPAccountPassword)) ) {
        Write-Host "Try another SPN Credential - registrationSpAccountName"
        $loginResult = az login --service-principal -u $registrationSpAccountName -p $registrationSPAccountPassword --tenant $spTenant | ConvertTo-Json
    }
    
    if ($null -eq $loginResult) {
        Write-Host "Try another SPN Credential - spAppId"
        $loginResult = az login --service-principal -u $spAppId -p $spPassword --tenant $spTenant | ConvertTo-Json
    }

    if ($null -eq $loginResult) {
        throw "Cannot do the az login"
    }

    az account set --subscription $subscriptionId | Out-Null
    Write-Host "Logged with AZ Login, and set the subscription to $subscriptionId." 
}

function Get-HybridAksVersionInfo {
    param(
        [parameter(Mandatory = $false, HelpMessage = "Set to true if we need save the full information to file.")]
        [switch]$ExportToFile
    )

    $remoteRootBlock = {
        #ASZ version info
        $aszInfo = Get-AzurestackHci 
        $stampInfo = Get-StampInformation
    
        $aszInfo | Add-Member -NotePropertyName "DeploymentID" -NotePropertyValue $stampInfo.DeploymentID
        $aszInfo | Add-Member -NotePropertyName "StampVersion" -NotePropertyValue $stampInfo.StampVersion
        $aszInfo | Add-Member -NotePropertyName "DomainFQDN" -NotePropertyValue $stampInfo.DomainFQDN
    
        $extensionInfoString = az extension list
    
        $extensions = [string]::Join(" ", $extensionInfoString) | ConvertFrom-Json
    
        $componentInfo = @{}

        $getArbResult = Get-ArcHciMgmt
        
        if (($getArbResult -is [array]) -and $getArbResult[$getArbResult.count - 1] -is [HashTable]) {
            $componentInfo = $getArbResult[$getArbResult.count - 1]
        }
        elseif ($getArbResult -is [hashtable]) {
            # I think that Get-archcimgmt should return a hashtable, but in recent build, it will retur an array, it should be a bug.
            $componentInfo = $getArbResult
        }

        $componentInfo["MocBuild"] = Get-MocConfig
        $componentInfo["ASZInformation"] = $aszInfo
        $extensions | ForEach-Object { $componentInfo["cli-$($_.name)"] = $_ }

        return $componentInfo
    }
    
    $remote_session = Get-HostPSSession -DoAzlogin
    Write-Host "Connecting to $($remote_session.ComputerName) with user $($remoteCred.UserName) to collect version information."
    $result = Invoke-Command -Session $remote_session -ScriptBlock  $remoteRootBlock
    
    #ASZ Version
    Write-Host "ASZ: $($result['ASZInformation'].StampVersion)"
    
    # Arb Version
    Write-Host "ArcAppliace: $($result['cli-arcappliance'].version)"
    
    # Moc Version
    $mocConfig = $result["MocBuild"]
    Write-Host "MOC: $($mocConfig.version) (catalog: $($mocConfig.catalog); ring: $($mocConfig.ring))"
    
    #HybridAks Version
    $hybridAks = $result["HybridaksExtension"]
    if ($hybridAks.provisioningState -eq "Succeeded") {
        $hybridAksVersion = $hybridAks.version
        if ($true -eq [string]::IsNullOrEmpty($hybridAksVersion)) {
            $hybridAksVersion = $hybridAks.currentVersion
        }
        Write-Host "HybridAKSExtension: $($hybridAksVersion) (ReleaseTrain: $($hybridAks.releaseTrain))"
    }
    else {
        Write-Host "HybridAKSExtension: $($hybridAks.provisioningState)"
    }

    if ($ExportToFile.IsPresent) {
        $fileName = "VersionInfo_$((Get-Date).ToString(`"MMdd-HHmmss`")).json"
        $filePath = Join-Path $PSScriptRoot $fileName
        New-Item -Path $filePath 
        $result | ConvertTo-Json | Add-Content -Path $filePath
        Write-Host "The full version info exported to file $filePath"

        Start-Process 'C:\WINDOWS\system32\notepad.exe' $filePath
    }
}

function Get-HostPSSession {
    param(

        [parameter(Mandatory = $false, HelpMessage = "Set the index of host.")]
        [int]$HostIndex = 0,
        [parameter(Mandatory = $false, HelpMessage = "Set to true to do the az login.")]
        [switch]$DoAzlogin
    )
    Write-Host "ASZ Host Lists: $([string]::Join(',', $hostLists))"

    if ($HostIndex -ge $hostLists.Count) {
        Write-Host "Specified HostIndex is $HostIndex, it is bigger than $($hostLists.Count -1), we will choose last machine"
        $targetHost = $hostLists[$hostLists.Count - 1] 
    }
    else {
        $targetHost = $hostLists[$HostIndex] 
    }

    Write-Host "Connecting to $targetHost with user $($remoteCred.UserName) "
    $remoteSession = New-PSSession -ComputerName $targetHost -Credential $remoteCred -Authentication Credssp

    if ($DoAzlogin.IsPresent) {
        $azloginScript = {
            if (!(($env:PATH).Split(';') -contains $azCliPath)) {
                $env:PATH += ";C:\Program Files (x86)\Microsoft SDKs\Azure\CLI2\wbin;"    
            }

            if (($false -eq [string]::IsNullOrEmpty($using:registrationAccountName)) -and ($false -eq [string]::IsNullOrEmpty($using:registrationAccountPassword))) {
                if ($false -eq ($using:registrationAccountName).contains("@")) {

                    Write-Host "Try to logon with spn"
                    $loginResult = az login --service-principal -u $using:registrationAccountName -p $using:registrationAccountPassword --tenant $using:spTenant | ConvertTo-Json
                }
                else {
                    Write-Host "Try to logon with username/password"
                    $loginResult = az login -u $using:registrationAccountName -p $using:registrationAccountPassword --tenant $using:spTenant | ConvertTo-Json
                }
            }

            if (($null -eq $loginResult) -and ($false -eq [string]::IsNullOrEmpty($using:registrationSpAccountName)) -and ($false -eq [string]::IsNullOrEmpty($using:registrationSPAccountPassword))) {
                Write-Host "Try another SPN Credential registrationSpAccountName"
                # got an error with login with account and password, try to login with spn directly.
                $loginResult = az login --service-principal -u $using:registrationSpAccountName -p $using:registrationSPAccountPassword --tenant $using:spTenant | ConvertTo-Json
            }

            if ($null -eq $loginResult) {
                Write-Host "Try to logon with another SPN -SpAppId"
                # got an error with login with account and password, try to login with spn directly.
                $loginResult = az login --service-principal -u $using:spAppId -p $using:spPassword --tenant $using:spTenant | ConvertTo-Json
            }
        
            if ($null -eq $loginResult) {
                throw "Cannot do the az login"
            }
        
            az account set --subscription $using:subscriptionId | Out-Null
            Write-Host "Logged with AZ Login, and set the subscription to $using:subscriptionId." 
        }

        Write-Host "DoAzlogin specified, loging in in the remote session"

        Invoke-Command  -ScriptBlock  $azloginScript -Session $remoteSession
    }

    return $remoteSession
}

function Enter-HostPSSession {
    param(
        [parameter(Mandatory = $false, HelpMessage = "Set the index of host.")]
        [int]$HostIndex = 0,
        [parameter(Mandatory = $false, HelpMessage = "Set to true to do the az login.")]
        [switch]$DoAzlogin
    )

    $remoteSession = Get-HostPSSession -HostIndex $HostIndex -DoAzlogin:$DoAzlogin
    Write-Host "Enter remote PSSession ..."

    Enter-PSSession -Session $remoteSession
}

function Get-ARBLogsFromHost {
    param(
        [parameter(Mandatory = $false, HelpMessage = "Set a local folder to save collected logs, if not set, save the logs to script root folder.")]
        [string]$LogDir
    )

    $LogName = "$domain-ArcHciLogs-$((Get-Date).ToString(`"MMdd_HHmm`")).zip"
    $remoteRootBlock = {
        $logResult = Get-ArcHciLogs

        $remoteLogPath = $logResult[$logResult.count - 1]

        return $remoteLogPath
    }
    
    $remote_session = Get-HostPSSession -DoAzlogin

    Write-Host "Connecting to Connecting to $($remote_session.ComputerName) to collect ARB logs"
    $logResult = Invoke-Command  -ScriptBlock  $remoteRootBlock -Session $remote_session

    # The last line should be the log path.
    if ($logResult -is [array] ) {
        $remoteLogPath = $logResult[$logResult.count - 1]
    }
    else {
        $remoteLogPath = $logResult
    }

    if ([string]::IsNullOrEmpty($LogDir)) {
        $LogDir = $PSScriptRoot
    }
    elseif (!(Test-Path $LogDir)) {
        Write-Host "Cannot access $LogDir, will copy log to $PSScriptRoot"
        $LogDir = $PSScriptRoot
    }

    $logPath = Join-Path $LogDir $LogName
    Copy-Item -Path $remoteLogPath -Destination $logPath -FromSession $remote_session

    Write-Host "ARB log collected and copied to $logPath"
    Remove-PSSession -Session $remote_session
}

function Update-K8SExtension {
    param(

        [parameter(Mandatory = $false, HelpMessage = "Specify the Name of the extension instance, defaultValue is hybridaksextension.")]
        [string]$ExtensionName = "hybridaksextension",
        [parameter(Mandatory = $false, HelpMessage = "Specify the release train for the HybridAKS extension type, defaultValue is prerelease.")]
        [string]$NewReleaseTrain = "prerelease",
        [parameter(Mandatory = $false, HelpMessage = "Specify the version to install for the extension instance, null means latest version.")]
        [string]$NewVersion = ""
    )

    $remoteRootBlock = {
        $aszResoureUri = Get-AzurestackHci | Select-Object -ExpandProperty AzureResourceUri
        $regex = "/Subscriptions/(?<Subscription>.*)/resourceGroups/(?<ResourceGroup>.*)/providers/Microsoft.AzureStackHCI/clusters/(?<ClusterName>.*)"
        if ($aszResoureUri -match $regex) {
            $subscription = $matches["Subscription"]
            $resourceGroup = $matches["ResourceGroup"]
            $clusterName = $matches["ClusterName"]
        }
        else {
            throw "Failed to get AzureStackHCI info."
        }

        $arbName = "$clusterName-arcbridge"

        try {
            $extensionInfo = Invoke-AzCli -Script { az k8s-extension show --cluster-type appliances --name $using:ExtensionName --cluster-name $arbName --resource-group $resourceGroup --only-show-errors } -AsJson
            Write-Host "the original extension version is : [$($extensionInfo.releaseTrain): $($extensionInfo.currentVersion) $($extensionInfo.version)]"
        }
        catch {
            throw "Didn't find specified extension in $using:ExtensionName resource group $resourceGroup. "
        }

        Write-Host "Updating $using:ExtensionName with [$($using:NewReleaseTrain): $($using:NewVersion)]"
        if ($true -eq [string]::IsNullOrEmpty($NewVersion)) {
            Write-Host "Call command to update extension  az k8s-extension update --cluster-type appliances --name $using:ExtensionName --cluster-name $arbName --resource-group $resourceGroup --release-train $using:NewReleaseTrain"
            $updateResult = Invoke-AzCli -Script { az k8s-extension update --cluster-type appliances --name $using:ExtensionName --cluster-name $arbName --resource-group $resourceGroup --release-train $using:NewReleaseTrain --auto-upgrade false } -AsJson
        }
        else {
            Write-Host "Call command to update extension  az k8s-extension update --cluster-type appliances --name $using:ExtensionName --cluster-name $arbName --resource-group $resourceGroup --release-train $using:NewReleaseTrain  --version $using:NewVersion"
            $updateResult = Invoke-AzCli -Script { az k8s-extension update --cluster-type appliances --name $using:ExtensionName --cluster-name $arbName --resource-group $resourceGroup --release-train $using:NewReleaseTrain --version $using:NewVersion --auto-upgrade false } -AsJson
        }
        Write-Host "HybridAksExtension updated to  [$($updateResult.releaseTrain): $($updateResult.version)]"
    }
    
    $remote_session = Get-HostPSSession -DoAzlogin
    Import-LocalFunctionToRemoteSession -RemoteSession $remote_session -LocalFunctionName "Invoke-AzCli"
    Write-Host "Connecting to $($remote_session.ComputerName) to update extension"
    Invoke-Command  -ScriptBlock  $remoteRootBlock -Session $remote_session
}

function Get-ECELogs {
    Param
    (
        [parameter(Mandatory = $false)]
        [string]$ActionPlanInstanceId = "",
        [parameter(Mandatory = $false)]
        [string[]]$TargetMachines = @(),
        [parameter(Mandatory = $false)]
        [string]$logFolderPath = ""
    )
    
    if (($null -eq $TargetMachines) -or ($TargetMachines.Count -eq 0)) {
        $TargetMachines = $hostLists
        Write-Host "Getting EceLogs from all hosts $hostLists"
    }
    else {
        <# Action when all if and elseif conditions are false #>
        Write-Host "Getting EceLogs from $targetMachines"
    }

    $remoteCred = New-Credential -UserName $hostAccountName -Password $hostAccountPassword

    if ($true -eq [string]::IsNullOrEmpty($logFolderPath)) {
        $logFolderPath = Join-Path $PSScriptRoot "$($(Get-Date).tostring('MMdd-HHmmss'))"
    }

    if ($false -eq (Test-Path $logFolderPath -PathType Container)) {
        New-Item -Path $logFolderPath -ItemType Directory
    }

    Write-Host "The log will copy to $logFolderPath"

    foreach ($machine in $TargetMachines) {
        $driveLetter = "X"
        if ( $null -eq (Get-PSDrive -Name $driveLetter -ErrorAction Ignore)) {
            New-PSDrive -Name $driveLetter -PSProvider FileSystem -Root "\\$machine\c`$\Observability" -Credential $remoteCred -Persist
        }

        $machineECELogPath = Join-Path $logFolderPath "$machine"
        New-Item -Path $machineECELogPath -ItemType Directory

        Robocopy.exe "$($driveLetter):\ECEAgent" "$machineECELogPath\ECEAgent" /mir
        Robocopy.exe "$($driveLetter):\ECE" "$machineECELogPath\ECE" /mir
        Write-Host "Copied ECE logs from $machine to $machineECELogPath"
        Remove-PSDrive -Name $driveLetter
    }

    $resultFile = Join-Path $logFolderPath "result.log"
    New-Item -Path $resultFile -ItemType File
    $allEvents = @()
    # Copied all the files, include ETL file and ZIP file. if we are processing recent event, maybe we can ignore the ZIP file.
    Write-Host "Reading Event from all ETL files"
    $allEtlFiles = Get-ChildItem -Path $logFolderPath -Recurse -Filter "*.etl"
    $allEtlFiles | ForEach-Object { $allEvents += $(get-asEvent $_.FullName) }

    $targetEvents = $allEvents | Where-Object { $_.EventType -match "Instance(Verbose|Warning|Error)" } 
    if ($false -eq [string]::IsNullOrEmpty($ActionPlanInstanceId )) {
        $targetEvents = $targetEvents | Where-Object { $_.ActionPlanInstanceId -eq $ActionPlanInstanceId }
    }

    $targetEvents | Sort-Object TimeStamp | ForEach-Object { Add-Content -Path $resultFile -Value $_.FormattedMessage }

    Write-Host "Saved all related events to $resultFile"
}

function Launch-ManagementClusterSession {
    Param
    (
        [switch]$LaunchSshWindow
    )

    if ($false -eq (Test-Path ".\kubectl.exe")) {
        # download and install kubectl.exe
        curl.exe -LO "https://dl.k8s.io/release/v1.27.3/bin/windows/amd64/kubectl.exe"                                                                                                                                                                                                                                         
    }    

    $managementClusterFolder = Join-Path $env:USERPROFILE "ManagementCluster"

    Ensure-CleanFolder -folderPath $managementClusterFolder

    $sshPrivateKey = Join-Path $managementClusterFolder "sshlogkey"
    $sshCertificateFile = Join-Path $managementClusterFolder "sshlogkey-cert.pub"
    $kubeConfigPath = Join-Path $managementClusterFolder "kubeconfig"

    # Create the folder again.
    New-Item -Path $managementClusterFolder -ItemType Directory -ErrorAction SilentlyContinue

    # Copy the ssh key from host to local.
    $hostSession = Get-HostPSSession -DoAzlogin
    Import-LocalFunctionToRemoteSession -RemoteSession $hostSession -LocalFunctionName "Invoke-AzCli"
    $remoteScript = {
        $sshKeys = @{}
        $remoteSshkeyPath = "C:\ProgramData\kva\.ssh\logkey"
        # Check whether sshkey exist, if yes, get the content.
        if (Test-Path -Path $remoteSshkeyPath) {
            $sshKeyContent = Get-Content -Raw -Path $remoteSshkeyPath
            $sshKeys["ssh-privateKey"] = $sshKeyContent
        }
        else {
            # No ssh key, call Rest API to get the keys.
            $aszResoureUri = Get-AzurestackHci | Select-Object -ExpandProperty AzureResourceUri
            $regex = "/Subscriptions/(?<Subscription>.*)/resourceGroups/(?<ResourceGroup>.*)/providers/Microsoft.AzureStackHCI/clusters/(?<ClusterName>.*)"
            if ($aszResoureUri -match $regex) {
                $subscription = $matches["Subscription"]
                $resourceGroup = $matches["ResourceGroup"]
                $clusterName = $matches["ClusterName"]
            }
            $requestUri="https://management.azure.com/subscriptions/$subscription/resourceGroups/$resourceGroup/providers/Microsoft.ResourceConnector/appliances/$clusterName-arcbridge/listkeys?api-version=2022-10-27"
            
            $result = Invoke-AzCli "az rest --method post --url $requestUri" | ConvertFrom-Json 
            $sshKeys["kubeConfig"] = [System.Text.Encoding]::UTF8.GetString([Convert]::FromBase64String($result.kubeconfigs.value))
            $sshkeys["ssh-privateKey"] = [System.Text.Encoding]::UTF8.GetString([Convert]::FromBase64String($result.sshKeys.InternalManagementDevKey.privateKey))
            $sshkeys["ssh-certificateFile"] = [System.Text.Encoding]::UTF8.GetString([Convert]::FromBase64String($result.sshKeys.InternalManagementDevKey.certificate))
        }

        return $sshKeys
    }

    $sshkeys = Invoke-Command -ScriptBlock  $remoteScript -Session $hostSession
    if($sshkeys.contains("ssh-privateKey")){
        $sshPrivateKeyContent = $sshkeys["ssh-privateKey"]
        Set-Content -Path $sshPrivateKey -Value $sshPrivateKeyContent -NoNewline
    }

    if($sshkeys.contains("ssh-certificateFile")){
        $sshCertificateContent = $sshkeys["ssh-certificateFile"]
        Set-Content -Path $sshCertificateFile -Value $sshCertificateContent -NoNewline
    }
    
    # Get the control plane IP
    $controlPlaneIp = Invoke-Command -Session $hostSession -ScriptBlock { return (Get-ArcHciConfig).controlPlaneIP }

    # call copy the management cluster admin.conf
    write-host "Run command to set key: "
    if (Test-Path -Path $sshCertificateFile) {
        SSH clouduser@$controlplaneIP -i $sshPrivateKey -o StrictHostKeyChecking=no -o CertificateFile="$sshCertificateFile" "sudo cp /etc/kubernetes/admin.conf ~/ ; sudo chmod 666 ~/admin.conf"
        # Copy items to documents folder.
        scp -o "StrictHostKeyChecking=accept-new"  -o CertificateFile="$sshCertificateFile" -i $sshPrivateKey clouduser@$controlplaneIP`:~/admin.conf $kubeConfigPath
        }
    else {
        SSH clouduser@$controlplaneIP -i $sshPrivateKey -o StrictHostKeyChecking=no  "sudo cp /etc/kubernetes/admin.conf ~/ ; sudo chmod 666 ~/admin.conf"
        # Copy items to documents folder.
        scp -o "StrictHostKeyChecking=accept-new" -i $sshPrivateKey clouduser@$controlplaneIP`:~/admin.conf $kubeConfigPath
    }

    $env:kubeconfig = $kubeConfigPath
    Write-Host "Kubectl downloaded and configured."
    # set the kubeconfig path so that kubectl get use it directly.
    if ( $LaunchSshWindow.IsPresent) {
        LaunchSshRemote -SshKeyPath $sshPrivateKey -SshCertificatePath $sshCertificateFile -WindowTitle "ManagementCluster" -MachineIp $controlPlaneIp -KubeConfigPath $kubeConfigPath
    }
}

function Launch-TargetClusterSession {
    Param
    (
        [parameter(Mandatory = $true)]
        [string] $ClusterName,
        [parameter(Mandatory = $false)]
        [string] $ResourceGroup = "",
        [parameter(Mandatory = $false)]
        [string] $targetClusterSshkey = "",
        [switch] $LaunchSshWindow
    )

    # If caller didn't specify the ResourceGroup, try to get the resource group of HCI.
    if ($true -eq [string]::IsNullOrEmpty($ResourceGroup)) {
        Write-Host "didn't specify the resource group, try to get the HCI Resource group and use it."
        $remoteSession = Get-HostPSSession -DoAzlogin
        $hciArmId = Invoke-Command -Session $remoteSession -ScriptBlock { return (get-azurestackhci).AzureResourceUri }
        $splitArray = $hciArmId -split "/"
        #$subscriptionId = $splitArray[2] 
        $ResourceGroup = $splitArray[4]
        Write-Host "got the resource group: $ResourceGroup"
    }

    DoAzLogin 
    
    $clusterInfo = Invoke-AzCli -Script "az aksarc show --name $ClusterName -g $ResourceGroup --only-show-errors -o json" | ConvertFrom-Json
    
    if ($null -eq $clusterInfo) {
        Write-Host "Cannot find the cluster, you can call Create-SampleProvisionedCluster to create it." -ForegroundColor Red
        throw "Cannot get the AksArc Cluster from $ResourceGroup with name $ClusterName."
    }

    Write-Host "Got the clusterInfo of $ClusterName`:"
    Write-Host $clusterInfo

    # The cluster is exist, get the data.
    $targetClusterFolder = Join-Path $env:USERPROFILE "TargetCluster"
    # Create the folder again.
    New-Item -Path $targetClusterFolder -ItemType Directory -ErrorAction SilentlyContinue
    $sshKey = Join-Path $targetClusterFolder "logkey"
    $kubeConfigPath = Join-Path $targetClusterFolder "kubeconfig"

    Invoke-AzCli -Script "az aksarc get-credentials --name $ClusterName -g $resourceGroup --file $kubeConfigPath --admin"
    Write-Host "kubeconfig for $ClusterName saved in $kubeConfigPath"
    $env:kubeconfig = $kubeConfigPath
    # If no kubectl, try to download one.
    if ($false -eq (Test-Path ".\kubectl.exe" -PathType Leaf)) {
        curl.exe -LO "https://dl.k8s.io/release/v1.27.3/bin/windows/amd64/kubectl.exe"                                                                                                                                                                                                                                         
    }
    write-host "Kubectl downloaded and configured."

    if ($LaunchSshWindow.IsPresent) {
        if ($false -eq [string]::IsNullOrEmpty($targetClusterSshkey) -and 
            $true -eq ( Test-Path -Path $targetClusterSshkey)) {
            # the ssh key exist, 
            $sshKey = $targetClusterSshkey
        }
        else {
            $defaultSshKey = Join-Path $env:USERPROFILE ".ssh\id_rsa"
            Write-Host "SSH didn't specified, we will try to use the default key of $defaultSshKey"
            Write-Host "If you used Create-SampleProvisionedCluster to create cluster, the key will be generated at here"
            Copy-Item $defaultSshKey $sshKey
        }
        
        if ($false -eq ( Test-Path -Path $sshKey)) {
            throw "Cannot get the valid ssh key from $sshkey"
        }

        $nodes = .\kubectl.exe get nodes -l node-role.kubernetes.io/control-plane -o json | ConvertFrom-Json
        $firstNodeIP = $nodes.items[0].status.addresses[0].address

        LaunchSshRemote -SshKeyPath $sshKey -MachineIp $firstNodeIP -KubeConfigPath $kubeConfigPath -WindowTitle $ClusterName

    }
}

function LaunchSshRemote {
    Param
    (
        [parameter(Mandatory = $true)]
        [string]$SshKeyPath,
        [parameter(Mandatory = $true)]
        [string]$MachineIp,
        [parameter(Mandatory = $false)]
        [string]$SshCertificatePath = "",
        [parameter(Mandatory = $false)]
        [string]$KubeConfigPath,
        [parameter(Mandatory = $false)]
        [string]$WindowTitle = "SSH",
        [parameter(Mandatory = $false)]
        [string]$LogonUser = "clouduser"
    )

     $options = ""
    # if kubeConfigPath specified, set the kubeconfig to the remote machine.
    if ($false -eq [string]::IsNullOrEmpty($KubeConfigPath) -and 
    ($true -eq (Test-Path $KubeConfigPath -PathType Leaf))) {
        $copyKubeConfigCmd = "scp -o `"StrictHostKeyChecking=accept-new`"  -i $SshKeyPath $KubeConfigPath $LogonUser@$MachineIp`:~/config"
        Write-Host "Run command [$copyKubeConfigCmd] to copy kubeconfig."
        Invoke-Expression($copyKubeConfigCmd)
        $remoteCommand = "mkdir -p ~/.kube && sudo cp ~/config ~/.kube/config && sudo chmod 666 ~/.kube/config"
        Write-Host "Run command [$remoteCommand] in remote machine to set the kube config."

        if(Test-Path -Path $SshCertificatePath){
            $options = "-o CertificateFile=$SshCertificatePath"    
        }
        # else {
        #     ssh -o "StrictHostKeyChecking=accept-new" $LogonUser@$MachineIp -i $SshKeyPath $remoteCommand
        # }
        ssh -o "StrictHostKeyChecking=accept-new" $options -o "CertificateFile=$SshCertificatePath" $LogonUser@$MachineIp -i $SshKeyPath $remoteCommand

    }
    
    # open SSH to connect to remote and update the window title.
    $cmdcmd = "/K ssh -o `"StrictHostKeyChecking=accept-new`" $LogonUser@$MachineIp -t -i $SshKeyPath $options `"echo -ne '\033]0;$WindowTitle-$MachineIp\007';bash`""

    Write-Host "call [$cmdCmd] to start a new cmd window"

    Start-Process cmd.exe -ArgumentList $cmdCmd 

    # Copy the kubeConfig to remote machine.
    Write-Host "To copy files, call command: "
    Write-Host " `t VM to local: scp -i $SshKeyPath  $options $LogonUser@$MachineIp`:<source> <target>"
    Write-Host " `t local to VM: scp -i $SshKeyPath  $options <source> $LogonUser@$MachineIp`:<target> "
}

function Create-SampleProvisionedCluster {
    param (
        [parameter(Mandatory = $false)]
        [string] $ClusterName = "",
        [parameter(Mandatory = $false)]
        [string] $TargetVersion = ""
    )
    
    if ([string]::IsNullOrWhiteSpace($ClusterName) -eq $true) {
        $ClusterName = "spc-$($env:USERNAME)"
    }

    $configFile = "c:\asz-4-nodes.json"
    if ($false -eq (Test-Path -Path $configFile)) {
        $envName = $($env:COMPUTERNAME).Split('-')[0]
        $configFile = "c:\$envName.json"

        if ($false -eq (Test-Path -Path $configFile)) {
            $configFile = Get-ChildItem -Path "c:\" -Filter "*.json" | Where-Object { $_.Name -notlike "*full.json" } | Select-Object -First 1 -ExpandProperty FullName
        }
    }

    Write-Host "Check the environment config from $configFile"

    if ($false -eq (Test-Path -Path $configFile)) {
        throw "Cannot find the config file from $configFile"
    }

    $envJSONConfig = Get-Content -Path $configFile | ConvertFrom-Json
    $tenantNetConfig = $($envJSONConfig.ScaleUnits[0].DeploymentData.TenantNetworks | Where-Object { $_.NetworkType -eq 'Tenant' })[0]
    # Get first subnet allocated
    $tenantSubnetConfig = $($tenantNetConfig.Subnets)[0]
    # Get first IP pool
    $netIPPoolConfig = $($tenantSubnetConfig.IPPools)[0]
    $ipPoolStart = $netIPPoolConfig.StartingAddress
    $ipPoolEnd = $netIPPoolConfig.EndingAddress
    $ipAddressPrefix = $tenantSubnetConfig.AddressPrefix
    $gateway = $tenantSubnetConfig.DefaultGateway
    $dnsServer = $envJSONConfig.ScaleUnits.DeploymentData.InfrastructureNetwork.DnsServers
    $vlanId = $tenantSubnetConfig.VlanId

    $remoteSession = Get-HostPSSession
    $hciCluster = Invoke-Command -Session $remoteSession -ScriptBlock { return get-azurestackhci }
    $hciArmId = $hciCluster.AzureResourceUri
    $splitArray = $hciArmId -split "/"
    $subscriptionId = $splitArray[2] 
    $resourceGroup = $splitArray[4]
    $hciClusterName = $splitArray[8]
    $location = $hciCluster.Region

    $aadId = "f2b51c0a-5234-4963-89cb-5e1086724ac9"

    DoAzLogin
    
    $customLocationName = "$hciClusterName-customlocation"
    $customLocationId = Invoke-AzCli -script "az customlocation show -g $resourceGroup -n $customLocationName --query id -o tsv"
    if ([string]::IsNullOrEmpty($customLocationId)) {
        Write-Host "custom location id is empty. Stop running the script. Exit."
        throw
    }
        

    $endIp = [System.Net.IPAddress]::Parse($ipPoolEnd)

    # Reserve 5 IPs for create cluster control plane . 
    $endIpchanged = $ipPoolEnd
    $reservedIps = @()
    for ($i = 0; $i -lt 5; $i++) {
        $ip = [System.Net.IPAddress]::new($endIp.Address - $i * [Math]::Pow(2, 24))
        $reservedIps += $ip.ToString()
    }

    $endIpchanged = [System.Net.IPAddress]::new($endIp.Address - 5 * [Math]::Pow(2, 24)).ToString()
    Write-Host "Change the vnet node pool end ip to : $endIpchanged"
    Write-Host "Reserved IPs of $([string]::Join(', ', $reservedIps ))"

    # check whether the lnet already created, if not, create one.
    $lnets = Invoke-AzCli -Script "az stack-hci-vm network lnet list --resource-group $resourceGroup -o json --only-show-errors" | ConvertFrom-Json

    if ($lnets.count -ge 1) {
        $lnetId = $lnets[0].id
        Write-Host "Got the lnet,  select the first one [$lnetId] to create provisioned cluster"
    }
    else {
        Write-Host "Didn't get any lnet, create a new one."
        if ($hciClusterName.Length -lt 14) {
            $lnetPrefix = $hciClusterName
        }
        else {
            $lnetPrefix = $hciClusterName.substring(0, 14)
        }
        $lnetName = "$lnetPrefix-target-lnet"
        $vmSwitchName = Invoke-Command -Session $remoteSession -ScriptBlock { Get-VMSwitch  | Select-Object -ExpandProperty Name } 
        
        $createLnetScript = "az stack-hci-vm network lnet create --subscription $subscriptionId --resource-group $resourceGroup --custom-location $customLocationId --location $location --name $lnetName --ip-allocation-method Static --address-prefix $ipAddressPrefix --dns-servers $dnsServer --gateway $gateway --ip-pool-start $ipPoolStart --ip-pool-end $endIpchanged --vlan $vlanId --vm-switch-name '`"$vmSwitchName`"'"
        Write-Host "Create lnet with command [$createLnetScript]"
        Invoke-AzCli -Script $createLnetScript

        # lnet created, try to get lnet again
        $lnets = Invoke-AzCli -Script "az stack-hci-vm network lnet list --resource-group $resourceGroup -o json --only-show-errors" | ConvertFrom-Json
        $lnetId = $lnets[0].id
        Write-Host "Lnet created, Select the first one [$lnetId] to create provisioned cluster"
    }

    $createdCluster = Invoke-AzCli -Script "az aksarc show --name $ClusterName -g $resourceGroup --only-show-errors -o json" | ConvertFrom-Json

    if ($null -ne $createdCluster) {
        Write-Host "Cluster named $ClusterName already created in $resourceGroup"
    }
    else {
        Write-Host "Cluster named $ClusterName didn't create, try to create it now."
        
        # Get the supported version.
        if ([string]::IsNullOrEmpty($TargetVersion )) {
            Write-Host "Didn't specify the TargetVersion, query and random pick a supported version"
            $getVersionScript = "az aksarc get-versions --resource-group $resourceGroup --custom-location $customLocationId --only-show-errors"
            Write-Host "call command [$getVersionScript] to get support versions."
            $versionContent = Invoke-AzCli -Script $getVersionScript -AsJson
            
            # Random select a target version to create provisioned cluster, we didn't check readiness here, as we assume the images are ready.
            $TargetVersion = $versionContent.properties.Values | Get-Random | Select-Object -ExpandProperty patchVersions | Get-Member -MemberType Properties | Select-Object -ExpandProperty Name | Get-Random
        }

        $controlPlaneIp = Get-Random $reservedIps

        Write-Host "Select $TargetVersion to create cluster."

        $createClusterScript = "az aksarc create --name $ClusterName --resource-group $resourceGroup --custom-location $customLocationId --vnet-ids $lnetId --kubernetes-version $TargetVersion --aad-admin-group-object-ids $aadId --generate-ssh-keys --control-plane-node-count 3 --node-count 2 --load-balancer-count 0 --control-plane-ip $controlPlaneIp" 
        Write-Host "Create cluster with command: [$createClusterScript]"
        Invoke-AzCli -Script $createClusterScript 
        
        $createdCluster = Invoke-AzCli -Script "az aksarc show --name $ClusterName -g $resourceGroup --only-show-errors -o json" | ConvertFrom-Json

        if ($null -ne $createdCluster) {
            Write-Host "Cluster named $ClusterName already created in $resourceGroup"
        }
        else {
            throw "Failed to create cluster with name $ClusterName in $resourceGroup"
        }
    }

    # Get the Provisioned Cluster cred.
    Invoke-AzCli -Script "az aksarc get-credentials --name $ClusterName -g $resourceGroup --admin"
    Write-Host "kubeconfig for $ClusterName saved in $env:ComputerName"
}

function Import-LocalFunctionToRemoteSession {
    <#
    .DESCRIPTION
    This function imports a self-contained local function into a remote session.
    Self-contained means this function doesn't call any other functions that the remote session doesn't have,
    and this function doesn't use any variables in the local scope.
    All variables needed in the self-contained function should be passed in as parameters.
    #>
    param (
        [parameter(Mandatory = $true)]
        [System.Management.Automation.Runspaces.PSSession] $RemoteSession,

        [parameter(Mandatory = $true)]
        [ValidateNotNullOrEmpty()]
        [string] $LocalFunctionName
    )

    if ($null -eq $RemoteSession) {
        throw "Remote session cannot be null."
    }

    if ([string]::IsNullOrEmpty($LocalFunctionName)) {
        throw "LocalFunctionName cannot be empty."
    }
    
    $functionInstance = Get-Command -Name $LocalFunctionName
    if ($null -eq $functionInstance -or $null -eq $functionInstance.Definition) {
        throw "Didn't get the function definition of $LocalFunctionName"
    }

    $importBlock = { 
        param (
            [string]$functionBody,
            [string]$functionName
        )

        $functionDef = "function $functionName {$functionBody}"
        . ( [scriptblock]::Create($functionDef)) 
    }

    Invoke-Command -Session $RemoteSession -ScriptBlock $importBlock -ArgumentList $functionInstance.Definition, $LocalFunctionName

    Write-Host "Local function [$LocalFunctionName] imported to remote session on [$($RemoteSession.ComputerName)]"
}

Write-Host "This is a helper module for HybridAKS troubleshooting, it includes some below commands: "
Write-Host "Get-HybridAksVersionInfo            Print the ASZ, ARB, MOC, Extension Versions"
Write-Host "Get-HostPSSession                   Get a PsSession to Physical host, you can pick the different host with arguments"
Write-Host "Enter-HostPSSession                 Enter a PS Remote Session to Physical host, you can pick the different host with arguments"
Write-Host "Get-ARBLogsFromHost                 Collect the ARB logs and save to local or upload to blob"
Write-Host "Update-K8SExtension                 Update the Extension with specific version"
Write-Host "Get-ECELogs                         Collect the ECE logs, which be used to troubleshoot deployment/update issue, or other ECE actions."
Write-Host "Launch-ManagementClusterSession     Download Kubectl and ARB kubeconfig, and alternatively connect to the ARB VM."
Write-Host "Launch-TargetClusterSession         Download Kubectl and target cluster kubeConfig, and alternatively connect to the target cluster controlplane VM."
Write-Host "Create-SampleProvisionedCluster     Create a sample ProvisionedCluster on current HCI"

