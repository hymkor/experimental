Option Strict On
Option Infer On
Option Explicit On

Public Class Form1

    Private Sub ButtonX_Click(sender As Object, e As EventArgs)
        Dim b = TryCast(sender, Button)
        If b IsNot Nothing Then
            Me.TextBox1.Text = Me.TextBox1.Text & b.Text
        End If
    End Sub

    Private Sub Form1_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        For Each p As Control In Me.Controls
            Dim b = TryCast(p, Button)
            If b IsNot Nothing Then
                If p.Name.StartsWith("Button") Then
                    AddHandler b.Click, AddressOf Me.ButtonX_Click
                End If
            End If
        Next
    End Sub
End Class
