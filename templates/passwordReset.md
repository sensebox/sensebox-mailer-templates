---
language: de
fromName: openSenseMap
name: passwordReset
subject: Zurücksetzen deines Passworts
---

Hallo {{ .user.name }},

Jemand hat ein Zurücksetzen deines Passworts angefordert. Du kannst dein Passwort zurücksetzen, indem du auf den Link unten klickst.

[Passwort zurücksetzen]({{ .origin }}/account/password-reset?token={{ .token }})

Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:

{{ .origin }}/account/password-reset?token={{ .token }}

Dieser Link ist insgesamt 12 Stunden gültig.

Falls du kein neues Passwort angefordert hast, ignoriere diese E-Mail.

Dein Passwort wird sich nicht änden bevor du auf den Link geklickt hast und dort ein neues vergeben hast.

Wenn Du Fragen hast schreib uns eine Mail an: [support@sensebox.de](mailto:support@sensebox.de?Subject=Password%20Zur%C3%BCcksetzen%20f%C3%BCr%20{{ .user.email }})

Viele Grüße,

dein senseBox Team

---
language: en
fromName: openSenseMap
name: passwordReset
subject: Your password reset
---

Hi {{ .user.name }},

Someone has requested a link to change your password. You can do this by clicking the link below.

[Reset password]({{ .origin }}/account/password-reset?token={{ .token }})

If you are unable to click the link, you can also open this address with your web browser:

{{ .origin }}/account/password-reset?token={{ .token }}

This link is valid for 12 hours.

If you didn't request this, please ignore this email.

Your password won't change until you access the link above and create a new one.

If you have any questions, feel free to write us an email to: [support@sensebox.de](mailto:support@sensebox.de?Subject=Password%20Reset%20for%20{{ .user.email }})

Best wishes

your senseBox Team
