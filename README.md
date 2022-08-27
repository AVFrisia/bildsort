# bildsort

Bildsortierungsprogramm zur Archivierung

## Funktionsweise

bildsort wurde entwickelt, um unsortierte Bilder in Massen zu archivieren. Die Programmiersprache Go wird verwendet, um Abhängigkeiten statisch zu verknüpfen, was eine Cross-Kompilierung und eine direkte Ausführung in eingebetteten Umgebungen, wie z. B. auf einem NAS, ermöglicht.

### Sortierung

Alle Dateien im Quellordner werden auf Metadaten überprüft und dementsprechend nach Semester eingeordnet.

#### Caveat

Existierende Ordnerstrukturen gehen aktuell verloren, beispielsweise Unterordner pro Veranstaltung. Diese müssen bei Bedarf nachgetragen werden.

## Beispiel

```console
$ ./bildsort Unsortiert/ Bilder/
Sorting from Unsortiert to Bilder
2022/08/27 23:15:39 Unsortiert/IMG_3764.CR2 -> Bilder/SoSe 2022/IMG_3764.CR2
2022/08/27 23:15:39 Unsortiert/IMG_3765.CR2 -> Bilder/SoSe 2022/IMG_3765.CR2
2022/08/27 23:15:39 Unsortiert/IMG_3766.CR2 -> Bilder/WiSe 2022 - 2023/IMG_3766.CR2
2022/08/27 23:15:39 Unsortiert/IMG_3767.CR2 -> Bilder/WiSe 2022 - 2023/IMG_3767.CR2
```

## Cross-Kompilierung

Der Großteil der Endverbraucher-NAS (bsp. unser QNAP adH) basiert auf ARM-Prozessoren mit Linux. Hierfür kann das Programm mit `GOOS=linux GOARCH=arm go build .` von jedem normalen Rechner kompiliert werden.
