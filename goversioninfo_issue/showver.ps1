$v = [System.Diagnostics.FileVersionInfo]::GetVersionInfo($Args[0]) 
if( $v ){
    Write-Host("FileVersion=",$v.FileVersion)
    Write-Host("  FileMajorPart=",$v.FileMajorPart)
    Write-Host("  FileMinorPart=",$v.FileMinorPart)
    Write-Host("  FileBuildPart=",$v.FileBuildPart)
    Write-Host("  FilePrivatePart=",$v.FilePrivatePart)
    Write-Host("ProductVersion=",$v.ProductVersion)
    Write-Host("  ProductMajorPart=",$v.ProductMajorPart)
    Write-Host("  ProductMinorPart=",$v.ProductMinorPart)
    Write-Host("  ProductBuildPart=",$v.ProductBuildPart)
    Write-Host("  ProductPrivatePart=",$v.ProductPrivatePart)
}
