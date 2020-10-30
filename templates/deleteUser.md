---
language: de
fromName: openSenseMap
template: deleteUser
subject: Dein Account wurde gelöscht
---

<p>Hallo {{ .user.name }},</p>
<p>Dein Account und alle deine senseBoxen wurden gerade gelöscht. Schade, dass du dich gelöscht hast!<p>
<p>Dieser Vorgang ist endgültig. Wenn du gerne wieder teilnehmen möchtest, kannst du dich einfach auf <a href="https://opensensemap.org/" target="_blank">https://opensensemap.org/</a> neu registrieren.
<p>Wenn Du Fragen hast schreib uns eine Mail an: <a href="mailto:support@sensebox.de" target="_top">support@sensebox.de</a>.</p>
<p>Tschüss,<br />dein senseBox Team</p>

---
language: en
fromName: openSenseMap
template: deleteUser
subject: Your openSenseMap account has been deleted
---

<p>Dear {{ .user.name }},</p>
<p>Your account and all registered senseBoxes are being deleted. Sad to see you go!</p>
<p>This action is irreversible. If you want to participate again, register yourself a new account at <a href="https://opensensemap.org/" target="_blank">https://opensensemap.org/</a>.
<p>If you have any questions, feel free to write us an email to: <a href="mailto:support@sensebox.de" target="_top">support@sensebox.de</a>.</p>
<p>Best wishes<br />your senseBox Team</p>