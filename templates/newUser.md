---
language: de
fromName: openSenseMap
template: newUser
subject: Deine openSenseMap Registrierung
---

Hallo {{ .user.name }},

danke für deine Registrierung auf der openSenseMap! Du findest dein Profil unter [{{ .origin }}/account]({{ .origin }}/account). Hier kannst du neue senseBoxen anlegen und deine Daten bearbeiten.

Bitte bestätige außerdem deine E-Mail Adresse. Klicke einfach auf diesen [Link]({{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}).

Wenn sich der Link nicht anklicken lässt, kannst du auch diese Adresse mit deinem Browser öffnen:

{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}

Bitte antworte nicht direkt auf diese Mail. Falls du Fragen oder Anmerkungen hast schreibe uns gerne eine Nachricht an [support@sensebox.de](mailto:support@sensebox.de?Subject=Nutzer%20Registrierung%20{{ .email }}).

Dein

senseBox Team

---
language: en
fromName: openSenseMap
template: newUser
subject: Your openSenseMap registration
---

Hi {{ .user.name }},

thank you for registering yourself on the openSenseMap platform. You'll find your profile at [{{ .origin }}/account]({{ .origin }}/account). There, you can create new senseBoxes and change data and credentials.

Please confirm your email address. Just click on this [link]({{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}).

If you are unable to click the link, you can also open this address with your web browser:

{{ .origin }}/account/confirm-email?email={{ .email }}&token={{ .token }}

If you have questions or suggestions please do not answer on this email directly, but feel free to send a mail to [support@sensebox.de](mailto:support@sensebox.de?Subject=User%20Registration%20{{ .email }}).

Best wishes

your senseBox Team
