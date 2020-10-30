---
language: de
fromName: openSenseMap
template: passwordReset
subject: Zurücksetzen deines Passworts
---

<p>
  Hallo {{ .user.name }},
</p>
<p>
  Jemand hat ein Zurücksetzen deines Passworts angefordert. Du kannst dein Passwort zurücksetzen, indem du auf den Link unten klickst.
</p>
<p>
  <a href="{{ .origin }}/account/password-reset?token={{ .token }}" target="_blank">Passwort zurücksetzen</a>
  <br />
  Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:
  <br />
  {{ .origin }}/account/password-reset?token={{ .token }}
  <br />
  Dieser Link ist insgesamt 12 Stunden gültig.
</p>
<p>
  Falls du kein neues Passwort angefordert hast, ignoriere diese E-Mail.
</p>
<p>
  Dein Passwort wird sich nicht änden bevor du auf den Link geklickt hast und dort ein neues vergeben hast.
</p>
<p>
  Wenn Du Fragen hast schreib uns eine Mail an: <a href="mailto:support@sensebox.de?Subject=Password%20Zur%C3%BCcksetzen%20f%C3%BCr%20{{ .user.email }}" target="_top">support@sensebox.de</a>
</p>
<p>
  Viele Grüße,<br />dein senseBox Team
</p>

---
language: en
fromName: openSenseMap
template: passwordReset
subject: Your password reset
---

<p>
Hi {{ .user.name }},
</p>
<p>
Someone has requested a link to change your password. You can do this by clicking the link below.
</p>
<p>
<a href="{{ .origin }}/account/password-reset?token={{ .token }}" target="_blank">Reset password</a>
<br />
If you are unable to click the link, you can also open this address with your web browser:
<br />
{{ .origin }}/account/password-reset?token={{ .token }}
<br />
This link is valid for 12 hours.
</p>
<p>
If you didn't request this, please ignore this email.
</p>
<p>
Your password won't change until you access the link above and create a new one.
</p>
<p>
If you have any questions, feel free to write us an email to: <a href="mailto:support@sensebox.de?Subject=Password%20Reset%20for%20{{ .user.email }}" target="_top">support@sensebox.de</a>
</p>
<p>
Best wishes<br />your senseBox Team
</p>