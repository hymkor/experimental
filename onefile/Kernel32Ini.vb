Option Strict On
Option Infer On

Imports System.Runtime.InteropServices

Public Class Kernel32Ini
    Private Declare Auto Function GetPrivateProfileString_ Lib "kernel32.dll" _
        Alias "GetPrivateProfileString" ( _
        ByVal lpApplicationName As String, _
        ByVal lpKeyName As String, _
        ByVal lpDefault As String, _
        ByVal lpReturnedString As System.Text.StringBuilder, _
        ByVal nSize As UInt32, _
        ByVal lpFileName As String) As UInt32

    Public Shared Function GetValue(ByVal path As String, ByVal section As String, ByVal key As String, ByVal defaultValue As String) As String
        Dim buffer As New System.Text.StringBuilder(1000)
        GetPrivateProfileString_(section, key, defaultValue, buffer, CUInt(buffer.Capacity), path)
        Dim value As String = buffer.ToString()
        Return value
    End Function

    Declare Function GetPrivateProfileStringByByteArray Lib "kernel32.dll" _
    Alias "GetPrivateProfileStringA" _
    (ByVal lpAppName As String, ByVal lpKeyName As String,
     ByVal lpDefault As String, ByVal lpReturnedString As Byte(),
     ByVal nSize As Integer, ByVal lpFileName As String) As Integer


    Public Shared Function GetKeys(ByVal path As String, ByVal section As String) As String()
        Dim binbuf(1024) As Byte
        Dim size As Integer = GetPrivateProfileStringByByteArray( _
            section, Nothing, "", binbuf, binbuf.Length, path)
        If size <= 0 Then
            Return New String() {}
        End If
        Dim keys As String = System.Text.Encoding.Default.GetString(binbuf, 0, size - 1)
        Return keys.Split(Chr(0))
    End Function
End Class
