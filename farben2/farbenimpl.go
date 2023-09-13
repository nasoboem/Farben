package farben2

import . "zufallszahlen"


func Bunt () (r,g,b uint8){
	Randomisieren()
	grundfarbe := Zufallszahl(0,2)
	zweitfarbe := Zufallszahl(0,1)
	zweitfarbeWert := Zufallszahl (0,255)
	switch grundfarbe {
		case 0: 
		r=255
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
		}
		case 1:
		g=255
		if zweitfarbe ==0{
			r=uint8(zweitfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
		}
		case 2:
		b=255
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
		}else {
			r=uint8(zweitfarbeWert)
		}
	}
	return r, g, b
}

func Orange () (r,g,b uint8){
	Randomisieren()
	r= uint8(Zufallszahl(230,255))
	g= uint8(Zufallszahl(100,int64(r)))
	b= uint8(Zufallszahl(0,int64(g-100)))
	return r,g,b
}

func Rot () (r,g,b uint8){
	Randomisieren()
	r= 255
	g= uint8(Zufallszahl(0,100))
	b= uint8(Zufallszahl(0,100))
	return r,g,b
}

func S_W () (r,g,b uint8){
	Randomisieren()
	r= uint8(Zufallszahl(0,255))
	g= r
	b= r
	return r,g,b
}

func Alle () (r,g,b uint8){
	Randomisieren()
	r= uint8(Zufallszahl(0,255))
	g= uint8(Zufallszahl(0,255))
	b= uint8(Zufallszahl(0,255))
	return r,g,b
}

func Blau() (r,g,b uint8){
	Randomisieren()
	b= 255
	r= uint8(Zufallszahl(0,128))
	g= uint8(Zufallszahl(0,255))
	return r,g,b
}

func Gelb() (r,g,b uint8){
	Randomisieren()
	r= uint8(Zufallszahl(180,255))
	g= uint8(Zufallszahl(180,255))
	b= uint8(Zufallszahl(0,40))
	return r,g,b
}

func Grün() (r,g,b uint8){
	Randomisieren()
	g= uint8(Zufallszahl(160,255))
	r= uint8(Zufallszahl(0,int64(g-120)))
	b= uint8(Zufallszahl(0,int64(g-80)))
	return r,g,b
}

func Lila() (r,g,b uint8){
	b= uint8(Zufallszahl(230,255))
	r= uint8(Zufallszahl(90,int64(b)))
	g= uint8(Zufallszahl(0,80))
	return r,g,b
}

func Rose() (r,g,b uint8){
	var z int64
	z= Zufallszahl(0,100)
	if z<60{
		r= 255
		b= uint8(Zufallszahl(50,200))
		g= 0
	}else{
		g= uint8(Zufallszahl(128,255))
		r= uint8(Zufallszahl(0,20))
		b= uint8(Zufallszahl(0,20))
	}
	return r,g,b
}

func Weiß() (r,g,b uint8){
	r= uint8(Zufallszahl(150,255))
	g= uint8(Zufallszahl(150,255))
	b= uint8(Zufallszahl(150,255))
	return r,g,b
}

func Schwarz() (r,g,b uint8){
	r= uint8(Zufallszahl(0,110))
	g= uint8(Zufallszahl(0,110))
	b= uint8(Zufallszahl(0,110))
	return r,g,b
}

func Bunthell () (r,g,b uint8){
	Randomisieren()
	grundfarbe := Zufallszahl(0,2)
	zweitfarbe := Zufallszahl(0,1)
	zweitfarbeWert := Zufallszahl (0,255)
	drittfarbeWert := Zufallszahl (80, zweitfarbeWert)
	switch grundfarbe {
		case 0: 
		r=255
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
			b=uint8(drittfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
			g=uint8(drittfarbeWert)
		}
		case 1:
		g=255
		if zweitfarbe ==0{
			r=uint8(zweitfarbeWert)
			b=uint8(drittfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
			r=uint8(drittfarbeWert)
		}
		case 2:
		b=255
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
			r=uint8(drittfarbeWert)
		}else {
			r=uint8(zweitfarbeWert)
			g=uint8(drittfarbeWert)
		}
	}
	return r, g, b
}

func Buntdunkel () (r,g,b uint8){
	Randomisieren()
	grundfarbe := Zufallszahl(0,2)
	zweitfarbe := Zufallszahl(0,1)
	zweitfarbeWert := Zufallszahl (0,255)
	drittfarbeWert := Zufallszahl (0,zweitfarbeWert)
	switch grundfarbe {
		case 0: 
		r=0
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
			b=uint8(drittfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
			g=uint8(drittfarbeWert)
		}
		case 1:
		g=0
		if zweitfarbe ==0{
			r=uint8(zweitfarbeWert)
			b=uint8(drittfarbeWert)
		}else {
			b=uint8(zweitfarbeWert)
			r=uint8(drittfarbeWert)
		}
		case 2:
		b=0
		if zweitfarbe ==0{
			g=uint8(zweitfarbeWert)
			r=uint8(drittfarbeWert)
		}else {
			r=uint8(zweitfarbeWert)
			g=uint8(drittfarbeWert)
		}
	}
	return r, g, b
}

func Farbe1 () (r,g,b uint8){
	r= uint8(Zufallszahl(0,255))
	g= uint8(Zufallszahl(0,255))
	b= 128
	return r, g, b
}

func Grau () (r,g,b uint8){
	b= uint8(Zufallszahl(90,180))
	r= uint8(Zufallszahl(90,180))
	g= uint8(Zufallszahl(90,180))
	return r, g, b
}

