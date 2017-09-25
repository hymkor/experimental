
function Get-FieldAll($rs){
    for($i=0 ; $i -lt $rs.Fields.Count ; $i++ ){
        Write-Output $rs.Fields.Item($i).Value
    }
}

try{
    $access = New-Object -ComObject Access.Application
    $access.Visible = $true
    $mdbpath = (Resolve-Path "ìdê¸èîå≥.mdb")
    Write-Host $mdbpath
    $access.OpenCurrentDatabase($mdbpath)

    $db = $access.CurrentDb()
    $rs = $db.OpenRecordset("select * from [DENSEN]")

    $max=10
    while( -not $rs.Eof -and $max -gt 0 ){
        Write-Output (@(Get-FieldAll($rs)) -join ",")

        $max--
        $rs.MoveNext()
    }
}finally{
    if( $s -ne $null ){
        $rs.Close()
        $rs = $null
    }
    if( $access -ne $null ){
        $access.Quit()
        $access = $null
    }
}
