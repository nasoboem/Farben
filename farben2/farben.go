package farben2

type Farbe interface{
	Bunt()(r,g,b uint8)//leuchtende Farben
	S_W()(r,g,b uint8)//Graustufen
	Schwarz()(r,g,b uint8)//Dunkle Farben
	Weiß()(r,g,b uint8)//helle Farben
	Alle()(r,g,b uint8)//alle Farben
	//Farben ähnlich wie die Farbe im Funktionsnamen:
	Rot()(r,g,b uint8)
	Orange()(r,g,b uint8)
	Gelb()(r,g,b uint8)
	Grün()(r,g,b uint8)
	Blau()(r,g,b uint8)
	Lila()(r,g,b uint8)
	Grau()(r,g,b uint8)
	Bunthell()(r,g,b uint8)//helle Farben ohne Grau
	Buntdunkel()(r,g,b uint8)//dunkle Farben ohne Grau
	Regenbogen(r1, g1, b1, n uint8) (r,g,b uint8)//Zufällige Farbwerte 
								//zwischen Farbwert-n und Farbwert+n
}

