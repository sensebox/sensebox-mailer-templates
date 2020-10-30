---
language: de
fromName: openSenseMap
template: newBox
subject: Deine senseBox auf der openSenseMap
---

Hallo {{ .user.name }},

deine senseBox mit dem Namen "{{ .box.name }}" ist nun auf der openSenseMap angemeldet! 🎉 Vielen lieben Dank, dass du dich am Projekt beteiligst.

Du findest deine Station auf der openSenseMap unter dieser Adresse:<br /><a
      href="{{ .origin }}/explore/{{ .box._id }}" target="_blank">{{ .origin }}/explore/{{ .box._id }}</a></p>

Im Anhang befindet sich dein persönlicher **Arduino Sketch**. Falls du eine senseBox mit WiFi-Bee registriert hast, denke unbedingt daran dein WiFi-Netzwerknamen und das Passwort in den Arduino Sktech einzufügen, damit sich deine senseBox mit dem Internet verbinden kann. Eine Anleitung wie es damit weitergeht, findest du [**hier**](https://docs.sensebox.de/sensebox-home/home-schritt-1) in der Dokumentation.

Wenn Du Fragen hast schreib uns eine Mail an: <a href="mailto:support@sensebox.de?Subject=Hilfe%20bei%20der%20Einrichtungbody=Bitte%20bei%20jeder%20Anfrage%20die%20senseBox-ID%20({{ .box._id }})%20mit%20angeben.%20Danke!" target="_top">support@sensebox.de</a>.

Viel Spaß wünscht dein senseBox Team


---
language: en
fromName: openSenseMap
template: newBox
subject: Your new senseBox on openSenseMap
---

<p>Hello {{ .user.name }},</p>
<p>your senseBox "{{ .box.name }}" is now registered at openSenseMap! 🎉 Thank you very much for contributing!</p>
<p>You can view your senseBox at this location:<br /><a href="{{ .origin }}/explore/{{ .box._id }}" target="_blank">{{ .origin }}/explore/{{ .box._id }}</a></p>
<p>Please note your personal <strong>Arduino Sketch</strong> in the attachment of this mail. If you have registered a senseBox with WiFi-Bee make sure you set your WiFi credentials in the arduino sketch, so your senseBox can connect to the internet. You can find further instructions <strong><a href="https://en.docs.sensebox.de/sensebox-home/home-schritt-1/" target="_blank">here</a></strong> in the "First steps" of our senseBox:home book.</p>
<p>If you have any questions, feel free to write us an email to: <a href="mailto:support@sensebox.de?Subject=Help%20me%20with%20my%20senseBox&body=Please%20include%20your%20senseBox-ID%20({{ .box._id }})%20in%20every%20message.%20Thank you!" target="_top">support@sensebox.de</a>.</p>
<p>The senseBox team wishes you a lot of fun with your new senseBox</p>