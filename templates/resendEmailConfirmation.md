---
language: de
fromName: openSenseMap
template: resendEmailConfirmation
subject: Bestätigung deiner E-Mailadresse
---

Hallo {{ .user.name }},

Du hast bisher deine E-Mail Adresse noch nicht bestätigt. Dies kannst du tun, indem du auf den Link unten klickst.

[E-Mail bestätigen]({{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }})

Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:

{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}

Wenn Du Fragen hast schreib uns eine Mail an: [support@sensebox.de](mailto:support@sensebox.de?Subject=Email%20Best%C3%A4tigen%20f%C3%BCr%20{{ .email }})

Viele Grüße,

dein senseBox Team

---
language: en
fromName: openSenseMap
template: resendEmailConfirmation
subject: E-Mail address confirmation
---

Hi {{ .user.name }},

Your email address on {{ .origin }} isn't confirmed yet. You can do this by clicking the link below.

[Confirm email address]({{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }})

If you are unable to click the link, you can also open this address with your web browser:

{{ .origin }}/account/confirm-email?token={{ .token }}&email={{ .email }}

If you have any questions, feel free to write us an email to: [support@sensebox.de](mailto:support@sensebox.de?Subject=Email%20Confirmation%20for%20{{ .email }})

Best wishes

your senseBox Team
