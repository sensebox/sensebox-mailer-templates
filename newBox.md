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


<!-- ---
language: en
fromName: openSenseMap
template: newBox
subject: Your new senseBox on openSenseMap
--- -->