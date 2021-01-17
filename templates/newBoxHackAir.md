---
language: de
fromName: openSenseMap
name: newBoxHackAir
subject: Dein neues Gerät auf der openSenseMap
---

Hallo {{ .user.name }},

vielen Dank für die Registrierung deines hackAIR home v2 Feinstaubsensors "{{ .box.name }}" auf der openSenseMap!

🎉 Damit deine Daten auch die openSenseMap erreichen, musst du noch deinen Feinstaubsensor konfigurieren. Eine Anleitung findest du unter [https://docs.sensebox.de/opensensemap/opensensemap-hackair/](https://docs.sensebox.de/opensensemap/opensensemap-hackair/).

Vielen lieben Dank, dass du dich am Projekt beteiligst.

Deine senseBox-ID lautet: **{{ .box._id }}**

Du findest deine Station auf der openSenseMap unter dieser Adresse:

[{{ .origin }}/explore/{{ .box._id }}]({{ .origin }}/explore/{{ .box._id }})

Wenn Du Fragen hast schreib uns eine Mail an: [support@sensebox.de](mailto:support@sensebox.de?Subject=Hilfe%20bei%20der%20Einrichtung&body=Bitte%20bei%20jeder%20Anfrage%20die%20senseBox-ID%20({{ .box._id }})%20mit%20angeben.%20Danke!)

Viel Spaß wünscht dein senseBox Team

---
language: en
fromName: openSenseMap
name: newBoxHackAir
subject: Your device on openSenseMap
---

Hello {{ .user.name }},

Thank you for registering your hackAIR home v2 particulate matter sensor "{{ .box.name }}" on openSenseMap! 🎉 Now, you have to configure your device in order to submit measurements to the openSenseMap. You'll find instructions to do so [https://en.docs.sensebox.de/opensensemap/opensensemap-hackair/](https://en.docs.sensebox.de/opensensemap/opensensemap-hackair/). Thank you very much for contributing!

Your senseBox ID is: **{{ .box._id }}**

You can view your device at this location:

[{{ .origin }}/explore/{{ .box._id }}]({{ .origin }}/explore/{{ .box._id }})

If you have any questions, feel free to write us an email to: [support@sensebox.de](mailto:support@sensebox.de?Subject=Help%20me%20with%20my%20senseBox&body=Please%20include%20your%20senseBox-ID%20({{ .box._id }})%20in%20every%20message.%20Thank you!).

The senseBox team wishes you a lot of fun
