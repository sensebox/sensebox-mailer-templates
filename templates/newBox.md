---
language: de
fromName: openSenseMap
template: newBox
subject: Deine senseBox auf der openSenseMap
requireAttachment: true
---

Hallo {{ .user.name }},

deine senseBox mit dem Namen "{{ .box.name }}" ist nun auf der openSenseMap angemeldet! 🎉 Vielen lieben Dank, dass du dich am Projekt beteiligst.

Du findest deine Station auf der openSenseMap unter dieser Adresse:

[{{ .origin }}/explore/{{ .box._id }}]({{ .origin }}/explore/{{ .box._id }})

Im Anhang befindet sich dein persönlicher **Arduino Sketch**. Falls du eine senseBox mit WiFi-Bee registriert hast, denke unbedingt daran dein WiFi-Netzwerknamen und das Passwort in den Arduino Sktech einzufügen, damit sich deine senseBox mit dem Internet verbinden kann. Eine Anleitung wie es damit weitergeht, findest du [**hier**](https://docs.sensebox.de/sensebox-home/home-schritt-1) in der Dokumentation.

Wenn Du Fragen hast schreib uns eine Mail an: [support@sensebox.de](mailto:support@sensebox.de?Subject=Hilfe%20bei%20der%20Einrichtungbody=Bitte%20bei%20jeder%20Anfrage%20die%20senseBox-ID%20({{ .box._id }})%20mit%20angeben.%20Danke!).

Viel Spaß wünscht dein senseBox Team

---
language: en
fromName: openSenseMap
template: newBox
subject: Your new senseBox on openSenseMap
requireAttachment: true
---

Hello {{ .user.name }},

your senseBox "{{ .box.name }}" is now registered at openSenseMap! 🎉 Thank you very much for contributing!

You can view your senseBox at this location:

[{{ .origin }}/explore/{{ .box._id }}]({{ .origin }}/explore/{{ .box._id }})

Please note your personal **Arduino Sketch** in the attachment of this mail. If you have registered a senseBox with WiFi-Bee make sure you set your WiFi credentials in the arduino sketch, so your senseBox can connect to the internet. You can find further instructions [**hier**](https://en.docs.sensebox.de/sensebox-home/home-schritt-1) in the "First steps" of our senseBox:home book.

If you have any questions, feel free to write us an email to: [support@sensebox.de](mailto:support@sensebox.de?Subject=Help%20me%20with%20my%20senseBox&body=Please%20include%20your%20senseBox-ID%20({{ .box._id }})%20in%20every%20message.%20Thank you!).

The senseBox team wishes you a lot of fun with your new senseBox
