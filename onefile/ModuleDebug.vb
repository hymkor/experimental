Option Strict On
Option Infer On
Friend Module ModuleDebug
	Private _dbgWriter As New Lazy(Of System.IO.TextWriter)(
	Function() As System.IO.TextWriter
		Dim logPath As String = System.IO.Path.Combine(System.IO.Path.GetTempPath(), "DEBUG")
		If System.IO.Directory.Exists(logPath) Then
			Dim result = New System.IO.StreamWriter(System.IO.Path.Combine(logPath, My.Application.Info.AssemblyName & ".log"), False, System.Text.Encoding.Default)
			result.AutoFlush = True
			Return result
		Else
			Return System.IO.StreamWriter.Null
		End If
	End Function)

	Friend ReadOnly Property TempDebug As System.IO.TextWriter
		Get
			Return _dbgWriter.Value
		End Get
	End Property
End Module
