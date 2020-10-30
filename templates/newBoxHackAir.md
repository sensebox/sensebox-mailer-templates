---
language: de
fromName: openSenseMap
template: newBoxHackAir
subject: Dein neues Gerät auf der openSenseMap
---

<p>Hallo {{ .user.name }},</p>
<p>vielen Dank für die Registrierung deines hackAIR home v2 Feinstaubsensors "{{ .box.name }}" auf der openSenseMap!
  🎉 Damit deine Daten auch die openSenseMap erreichen, musst du noch deinen Feinstaubsensor konfigurieren. Eine
  Anleitung findest du unter <a href="https://docs.sensebox.de/opensensemap/opensensemap-hackair/"
    target="_blank">https://docs.sensebox.de/opensensemap/opensensemap-hackair/</a>.
  Vielen lieben Dank, dass du dich am Projekt beteiligst.</p>
<p>Deine senseBox-ID lautet: <b>{{ .box._id }}</b></p>
<p>Du findest deine Station auf der openSenseMap unter dieser Adresse:<br /><a
    href="{{ .origin }}/explore/{{ .box._id }}" target="_blank">{{ .origin }}/explore/{{ .box._id }}</a></p>
<p>Wenn Du Fragen hast schreib uns eine Mail an: <a
    href="mailto:support@sensebox.de?Subject=Hilfe%20bei%20der%20Einrichtung&body=Bitte%20bei%20jeder%20Anfrage%20die%20senseBox-ID%20({{ .box._id }})%20mit%20angeben.%20Danke!"
    target="_top">support@sensebox.de</a>.</p>
<p>Viel Spaß wünscht dein senseBox Team</p>

---
language: en
fromName: openSenseMap
template: newBoxHackAir
subject: Your device on openSenseMap
---

<p>Hello {{ .user.name }},</p>
<p>Thank you for registering your hackAIR home v2 particulate matter sensor "{{ .box.name }}" on openSenseMap! 🎉 Now, you have to configure your device in order to submit measurements to the openSenseMap. You'll find instructions (just in german) to do so <a href="https://en.docs.sensebox.de/opensensemap/opensensemap-hackair/" target="_blank">https://en.docs.sensebox.de/opensensemap/opensensemap-hackair/</a>. Thank you very much for contributing!</p>
<p>Your senseBox ID is: <b>{{ .box._id }}</b></p>
<p>You can view your device at this location:<br /><a href="{{ .origin }}/explore/{{ .box._id }}" target="_blank">{{ .origin }}/explore/{{ .box._id }}</a></p>
<p>If you have any questions, feel free to write us an email to: <a href="mailto:support@sensebox.de?Subject=Help%20me%20with%20my%20senseBox&body=Please%20include%20your%20senseBox-ID%20({{ .box._id }})%20in%20every%20message.%20Thank you!" target="_top">support@sensebox.de</a>.</p>
<p>The senseBox team wishes you a lot of fun</p>