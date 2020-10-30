---
language: de
fromName: openSenseMap
template: newSketch
subject: Dein neuer Sketch
---

<p>
Hallo {{ .user.name }},
</p>
<p>
  du hast gerade das Modell oder die Sensoren deiner Wetterstation auf der openSenseMap geändert. Deswegen schicken wir dir hier einen aktualisierten Arduino-Sketch, den du im Anhang findest.
</p>
<p><b>Achtung!</b></p>
<p>
  Der Sketch im Anhang befindet sich im Rohzustand. Das bedeutet:
</p>
<ul>
  <li>Netzwerkeinstellungen für Wifi oder statische Ethernetkonfigurationen müssen vom alten Sketch übernommen werden.</li>
<li>Jegliche Änderungen für zum Beispiel: andere Sensoren müssen vom alten Sketch übernommen werden.</li>
<li>Jede andere Änderung ist nicht enthalten!</li>
</ul>
<br />
<p>Deine senseBox-ID lautet: {{ .box._id }}</p>
<p>Außerdem wurden die folgenden IDs für die Sensoren deiner Station generiert:</p>
<ul>{{ range .box.sensors }}
    <li>{{ .title }} ({{ .sensorType }}): {{ ._id }}</li>
{{ end }}</ul>
<p>
  Wenn Du Fragen hast schreib uns eine Mail an: <a href="mailto:support@sensebox.de?Subject=Neuer%20Sketch" target="_top">support@sensebox.de</a>
</p>
<p>
  Viel Spaß wünscht dein senseBox Team
</p>

---
language: en
fromName: openSenseMap
template: newSketch
subject: Your new Sketch
---

<p>
  Hi {{ .user.name }},
</p>
<p>
  you've just changed your model or sensors of your weatherstation on the openSenseMap. You will find the new sketch in the attachment of this mail.
</p>
<p><b>Attention!</b></p>
<p>
  The in the attachment is in an untreated state. This means:
</p>
<ul>
  <li>Your network settings like wifi or ethernet configurations have to be taken over from your old sketch.</li>
<li>Every change you've made to include other sensors have to be taken over from your old sketch.</li>
<li>The sketch does not contain other changes you've made!</li>
</ul>
<br />
<p>Your senseBox ID is: {{ .box._id }}</p>
<p>In addition, the following IDs were generated for the sensors of your station:</p>
<ul>{{ range .box.sensors }}
    <li>{{ .title }} ({{ .sensorType }}): {{ ._id }}</li>
{{ end }}</ul>
<p>
  If you have any questions write us an email to: <a href="mailto:support@sensebox.de?Subject=New%20Sketch" target="_top">support@sensebox.de</a>
</p>
<p>
  The senseBox team wishes you a lot of fun with your new senseBox
</p>