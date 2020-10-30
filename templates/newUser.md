---
language: de
fromName: openSenseMap
template: newUser
subject: Deine openSenseMap Registrierung
---

<p>
  Hallo {{ .user.name }},
</p>
<p>
  danke für deine Registrierung auf der openSenseMap! Du findest dein Profil unter <a href="{{ .origin}}/account">{{ .origin}}/account</a>. Hier kannst du neue senseBoxen anlegen und deine Daten bearbeiten.
</p>
<p>
  Bitte bestätige außerdem deine E-Mail Adresse. Klicke einfach auf diesen <a href="{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}" target="_blank">Link</a>.
<br />
Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:
<br />
{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}
</p>
<p>
  Bitte antworte nicht direkt auf diese Mail. Falls du Fragen oder Anmerkungen hast schreibe uns gerne eine Nachricht an <a href="mailto:support@sensebox.de?Subject=Nutzer%20Registrierung%20{{ .email }}" target="_top">support@sensebox.de</a>.
</p>
<p>
  Dein<br />senseBox Team
</p>

---
language: en
fromName: openSenseMap
template: newUser
subject: Your openSenseMap registration
---

<p>
  Hi {{ .user.name }},
</p>
<p>
  thank you for registering yourself on the openSenseMap platform. You'll find your profile at  <a href="{{ .origin}}/account">{{ .origin}}/account</a>. There, you can create new senseBoxes and change data and credentials.
</p>
<p>
Please confirm your email address. Just click on this <a href="{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}" target="_blank">link</a>.
<br />
If you are unable to click the link, you can also open this address with your web browser:
<br />
{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}
</p>
<p>
  If you have questions or suggestions please do not answer on this email directly, but feel free to send a mail to <a href="mailto:support@sensebox.de?Subject=User%20Registration%20{{ .email }}" target="_top">support@sensebox.de</a>.
</p>
<p>
  Best wishes<br />your senseBox Team
</p>