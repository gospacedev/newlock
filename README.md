# newlock
[![MIT Licence](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/licenses/mit-license.php)

Newlock is a lightweight GUI application that generates truly random passwords easily

Donate Bitcoin: 1J2aMYUnkPXkzEYJUaZY5AKRGZU459NUrG

![Screenshot](https://user-images.githubusercontent.com/83633399/166413118-7d90a731-501d-447f-8f39-6babcde12184.png)

## Usage
Newlock uses the the Fyne library for it's GUI and sethvargo/go-password for generating the passwords

Install dependencies:
```
go get fyne.io/fyne/v2
go get github.com/sethvargo/go-password
```
To package:
```
fyne package -icon assets/orange-lock.png
```

And for specific versions:
```
fyne package -os linux -icon assets/orange-lock.png
fyne package -os windows -icon assets/orange-lock.png
```
