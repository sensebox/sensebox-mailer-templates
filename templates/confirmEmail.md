---
language: de
fromName: openSenseMap
template: confirmEmail
subject: openSenseMap Bestätigung deiner E-Mailadresse
---

<p>
  Hallo {{ .user.name }},
</p>
<p>
  Jemand hat eine Änderung deiner E-Mail Adresse angefordert. Dies kannst du tun, indem du auf den Link unten klickst.
</p>
<p>
  <a href="{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}" target="_blank">E-Mail bestätigen</a><br />
Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:
<br />
{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}
</p>
<p>
  Falls du keine Änderung deiner E-Mail Adresse angefordert hast, ignoriere diese E-Mail.
</p>
<p>
  Wenn Du Fragen hast schreib uns eine Mail an: <a href="mailto:support@sensebox.de?Subject=Email%20Best%C3%A4tigen%20f%C3%BCr%20{{ .email }}" target="_top">support@sensebox.de</a>
</p>
<p>
  Viele Grüße,<br />dein senseBox Team
</p>

---
language: en
fromName: openSenseMap
template: confirmEmail
subject: openSenseMap E-Mail address confirmation
---

<p>
  Hi {{ .user.name }},
</p>
<p>
Someone has requested to change your email address on {{ .origin }}. You can do this by clicking the link below.
</p>
<p>
  <a href="{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}" target="_blank">Confirm email address</a><br />
If you are unable to click the link, you can also open this address with your web browser:
<br />
{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}
</p>
<p>
  If you didn't request this, please ignore this email.
</p>
<p>
  If you have any questions, feel free to write us an email to: <a href="mailto:support@sensebox.de?Subject=Email%20Confirmation%20for%20{{ .email }}" target="_top">support@sensebox.de</a>
</p>
<p>
  Best wishes<br />your senseBox Team
</p>