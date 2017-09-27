function ToString(value){
    if( typeof value == "string" ){
        return "\"" + value + "\""
    }
    if( typeof value == "object" ){
        if( value.Value == null || value.Value == undefined ){
            return ""
        }
        var t = value.Type
        if( t == 10 ){ // string
            return "\"" + value.Value + "\""
        } else if( t == 1 ){ // boolean
            return value.Value + 0
        } else if( t == 8 ){ // date
            var dt = new Date(Date.parse( ""+value.Value ))
            return dt.getYear() + "/" +
                   (dt.getMonth()+1) + "/" +
                   dt.getDay() + " " +
                   dt.getHours() + ":" +
                   dt.getMinutes() + ":" +
                   dt.getSeconds()
        }
    }
    return ""+value
}

function DumpTable(db,tablename,outputf){
    var rs
    try{
        rs = db.OpenRecordset( "SELECT * FROM [" + tablename + "]" );
        if( ! rs.Eof ){
            var count = rs.Fields.Count
            var buffer = ""
            var comma = ""
            for(var i=0 ; i < count ; i++){
                buffer = buffer + comma + ToString(rs.Fields(i).Name)
                comma = ","
            }
            outputf( buffer )
            // WScript.Echo( buffer )
        }
        while( ! rs.Eof ){
            var count = rs.Fields.Count
            var buffer = ""
            var comma = ""
            for(var i=0 ; i < count ; i++){
                buffer = buffer + comma + ToString(rs(i))
                comma = ","
            }
            outputf( buffer )
            // WScript.Echo( buffer )
            rs.MoveNext();
        }
    }finally{
        if( rs != null ){
            rs.Close();
            rs = null;
        }
    }
}

var fsObj = new ActiveXObject("Scripting.FileSystemObject");
var access = new ActiveXObject("Access.Application");
try{
    access.Visible = true;
    var mdbPath = fsObj.GetAbsolutePathName(WScript.Arguments(0));
    access.OpenCurrentDatabase(mdbPath);
    var db = access.CurrentDb();
    for(var i=1 ; i < WScript.Arguments.length ; i++ ){
        var tableName = WScript.Arguments(i)
        var out = fsObj.OpenTextFile(tableName + ".csv",2 /* write */ ,true)
        try{
            DumpTable(db,tableName,function(line){
                out.WriteLine(line)
            })
        }finally{
            if( out != null ){
                out.Close()
            }
        }
    }
}finally{
    if( access != null ){
        access.CloseCurrentDatabase();
        access.Quit();
        access = null;
    }
}
