Module Module1
    Sub Main()
        Dim foo As String = Nothing
        Dim bar As String = CStr(foo)

        Console.Write("CStr(Nothing)={0}", If(bar, "(null)"))
        Console.ReadLine()
    End Sub
End Module
