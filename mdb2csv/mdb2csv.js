var fsObj = new ActiveXObject("Scripting.FileSystemObject");
var mdbPath = fsObj.GetAbsolutePathName(WScript.Arguments(0));
WScript.Echo("Path=" + mdbPath);
var access = new ActiveXObject("Access.Application");
var rs = null;
try{
    access.Visible = true;
    access.OpenCurrentDatabase(mdbPath);

    var db = access.CurrentDb();
    rs = db.OpenRecordset( "SELECT * FROM [DENSEN]" );
    while( ! rs.Eof ){
        WScript.Echo( rs("[����ID]")+","+rs("[����ID]")+","+rs("[�d�����d]") );
        rs.MoveNext();
    }
}finally{
    if( rs != null ){
        rs.Close();
        rs = null;
    }
    if( access != null ){
        access.CloseCurrentDatabase();
        access.Quit();
        access = null;
    }
}
