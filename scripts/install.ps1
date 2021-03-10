new-item -Path "C:\Program Files (x86)\" -Name "GhostDB" -ItemType "directory" -Force

Move-Item -Path ..\gdb.exe -Destination "C:\Program Files (x86)\GhostDB\gdb.exe"