go build
mkdir WorldEditor
mkdir WorldEditor\resource
copy WorldEditor.exe WorldEditor\
xcopy resource\ WorldEditor\resource\ /E
pause