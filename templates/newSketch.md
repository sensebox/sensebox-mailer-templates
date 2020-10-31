---
language: de
fromName: openSenseMap
template: newSketch
subject: Dein neuer Sketch
requireAttachment: true
---

Hallo {{ .user.name }},

du hast gerade das Modell oder die Sensoren deiner Wetterstation auf der openSenseMap geändert. Deswegen schicken wir dir hier einen aktualisierten Arduino-Sketch, den du im Anhang findest.

**Achtung!**

Der Sketch im Anhang befindet sich im Rohzustand. Das bedeutet:

- Netzwerkeinstellungen für Wifi oder statische Ethernetkonfigurationen müssen vom alten Sketch übernommen werden.
- Jegliche Änderungen für zum Beispiel: andere Sensoren müssen vom alten Sketch übernommen werden.
- Jede andere Änderung ist nicht enthalten!

Deine senseBox-ID lautet: **{{ .box._id }}**

Außerdem wurden die folgenden IDs für die Sensoren deiner Station generiert:

{{ range .box.sensors }}
- {{ .title }} ({{ .sensorType }}): {{ ._id }}
{{ end }}

Wenn Du Fragen hast schreib uns eine Mail an: [support@sensebox.de](mailto:support@sensebox.de?Subject=Neuer%20Sketch%20Box-Id%20{{ .box._id }})

Viel Spaß wünscht dein senseBox Team

---
language: en
fromName: openSenseMap
template: newSketch
subject: Your new Sketch
requireAttachment: true
---

Hi {{ .user.name }},

you've just changed your model or sensors of your weatherstation on the openSenseMap. You will find the new sketch in the attachment of this mail.

**Attention!**

The sketch in the attachment is in an untreated state. This means:

- Your network settings like wifi or ethernet configurations have to be taken over from your old sketch.
- Every change you've made to include other sensors have to be taken over from your old sketch.
- The sketch does not contain other changes you've made!

Your senseBox ID is: **{{ .box._id }}**

In addition, the following IDs were generated for the sensors of your station:

{{ range .box.sensors }}
- {{ .title }} ({{ .sensorType }}): {{ ._id }}
{{ end }}

If you have any questions write us an email to: [support@sensebox.de](mailto:support@sensebox.de?Subject=New%20Sketch%20Please%20Include%20box-id%20{{ .box._id }})

The senseBox team wishes you a lot of fun with your new senseBox
