function Get-FieldAll($rs){
    $fields = $rs.Fields
    $count = $Fields.Count
    $buffer = New-Object System.Text.StringBuilder
    for($i=0 ; $i -lt $count ; $i++ ){
        if( $i -ne 0 ){
            [void]$buffer.Append(",")
        }
        [void]$buffer.Append($fields.Item($i).Value)
    }
    return $buffer.ToString()
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
    while( -not $rs.Eof ){
        Write-Output (Get-FieldAll $rs)
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
