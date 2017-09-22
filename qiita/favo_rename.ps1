$fromName = "\\oldserver\"
$toName = "\\\newserver\"

$wsh = New-Object -ComObject WScript.Shell
Get-ChildItem (Join-Path $env:userprofile "Links") | %{
    if( $_.Extension -eq ".lnk" ){
        $shortcut = $wsh.CreateShortcut($_.FullName)
        $target = $shortcut.TargetPath
        if( $target.StartsWith($fromName) ){
            $newTarget = $toName + $target.substring($fromName.Length)
            if( Test-Path $newTarget ){
                Write-Output ("    " + $target)
                Write-Output ("--> " + $newTarget)
                $shortcut.TargetPath = $newTarget
                $shortcut.Save()
            }
        }
    }
}
