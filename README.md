# sensebox-mailer-templates

> ℹ️ This repository contains the old openSenseMap mailer templates and is deprecated and archived. New mailer templates can be found here: https://github.com/openSenseMap/mailer

Dieses Repository enthält die E-Mail Templates für den [sensebox-mailer](https://github.com/sensebox/sensebox-mailer)

## Templates

Templates liegen im [./templates](./templates) Verzeichnis und werden in [Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown) geschrieben.

Die einzelnen Templates liegen jeweils in Deutsch und Englisch vor. Beide Sprachen werden zusammen in einer Datei abgelegt.

Jedes Template beginnt mit einer so genannten "Frontmatter", welche Metadaten des Templates enthält. Diese beginnt und endet mit einer Zeile, welche nur `---` enthält.

| Feld | Erklärung |
| --- | --- |
| `name`     | Name des Templates wie er vom Mailer genannt wird. (Feld `template` im mailer request payload) |
| `language` | Sprache des Templates. Entweder `de` oder `en`. |
| `subject`  | Der Betreff der Mail. |
| `fromName` | Der Name des Absenders. |

Die Felder und deren Werte in der "Frontmatter" werden jeweils mit einem Doppelpunkt getrennt. Das Frontmatter sollte als YAML formatiert sein.

Die jeweils deutsche oder englische Version eines Templates mit gleichem `name` Feld befindet sich in einer Datei. Dabei folgt eine Datei immer dem Schema:

```
---
language: de
fromName: openSenseMap
name: newBoxHackAir
subject: Dein neues Gerät auf der openSenseMap
---

Hallo {{ .user.name }},

...

---
language: en
fromName: openSenseMap
name: newBoxHackAir
subject: Your device on openSenseMap
---

Hello {{ .user.name }},

...

```
