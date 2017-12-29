package main

var tree02id = `tree02`

var tree02name = `Dept. IM`
var tree02cname = `資訊管理學系`

var tree02desc = `This tree is for NTU IM`

var tree02svg = `
<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: Adobe Illustrator 22.0.1, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
    	 viewBox="0 0 755.2 800" enable-background="new 0 0 755.2 800" xml:space="preserve">
    <font horiz-adv-x="1000">
        <!-- Scada is a trademark of Jovanny Lemonad. -->
        <!-- Copyright: Copyright 2017 Adobe System Incorporated. All rights reserved. -->
        <font-face font-family="Scada-Regular" units-per-em="1000" underline-position="-75" underline-thickness="50"/>
        <missing-glyph horiz-adv-x="270"/>
        <glyph unicode=" " horiz-adv-x="290"/>
        <glyph unicode="A" horiz-adv-x="570" d="M15,0l205,700l130,0l205,-700l-100,0l-55,180l-230,0l-55,-180M380,250l-80,275C297,532 295,541 294,551C292,560 291,567 290,572l-3,23l-4,0l-3,-24C275,548 272,533 270,525l-80,-275z"/>
        <glyph unicode="B" horiz-adv-x="575" d="M90,700l190,0C350,700 404,684 443,652C481,619 500,577 500,525C500,472 481,430 442,399C424,386 405,376 385,370C410,364 433,353 455,338C502,302 525,253 525,190C525,131 505,84 465,51C424,17 363,0 280,0l-190,0M280,400C318,400 348,410 369,431C390,451 400,479 400,515C400,551 390,579 369,600C348,620 318,630 280,630l-90,0l0,-230M280,70C377,70 425,113 425,200C425,287 377,330 280,330l-90,0l0,-260z"/>
        <glyph unicode="C" horiz-adv-x="515" d="M55,350C55,475 77,566 121,624C165,681 226,710 305,710C352,710 395,701 434,682C453,673 468,664 480,655l-35,-60l-10,0C425,602 414,610 401,617C372,632 343,640 315,640C208,640 155,543 155,350C155,246 170,172 200,127C229,82 271,60 325,60C354,60 383,67 411,82C425,90 436,98 445,105l10,0l35,-60C479,36 463,26 444,17C407,-1 364,-10 315,-10C232,-10 168,19 123,76C78,133 55,225 55,350z"/>
        <glyph unicode="D" horiz-adv-x="585" d="M90,700l150,0C337,700 410,672 458,617C506,562 530,473 530,350C530,227 506,138 458,83C410,28 337,0 240,0l-150,0M240,70C303,70 351,91 383,134C414,177 430,249 430,350C430,451 414,523 383,566C351,609 303,630 240,630l-50,0l0,-560z"/>
        <glyph unicode="E" horiz-adv-x="510" d="M90,0l0,700l390,0l0,-70l-290,0l0,-230l230,0l0,-70l-230,0l0,-260l290,0l0,-70z"/>
        <glyph unicode="F" horiz-adv-x="500" d="M90,0l0,700l390,0l0,-70l-290,0l0,-240l230,0l0,-70l-230,0l0,-320z"/>
        <glyph unicode="G" horiz-adv-x="565" d="M55,350C55,431 66,498 88,552C109,606 139,646 177,672C214,697 257,710 305,710C352,710 395,701 434,682C453,673 468,664 480,655l-35,-60l-10,0C425,602 414,610 401,617C372,632 343,640 315,640C208,640 155,543 155,350C155,246 170,172 200,127C229,82 271,60 325,60C337,60 349,62 360,65C371,68 380,71 387,74C394,77 398,79 400,80l0,230l-115,0l0,70l185,0l30,-30l0,-305C487,34 470,25 451,17C412,-1 367,-10 315,-10C236,-10 173,20 126,80C79,139 55,229 55,350z"/>
        <glyph unicode="H" horiz-adv-x="630" d="M90,0l0,700l100,0l0,-310l250,0l0,310l100,0l0,-700l-100,0l0,320l-250,0l0,-320z"/>
        <glyph unicode="I" horiz-adv-x="280" d="M90,0l0,700l100,0l0,-700z"/>
        <glyph unicode="L" horiz-adv-x="490" d="M90,0l0,700l100,0l0,-630l290,0l0,-70z"/>
        <glyph unicode="M" horiz-adv-x="840" d="M90,700l100,0l210,-425C405,266 410,256 413,244C416,234 417,228 418,225l4,0C423,228 424,234 427,244C430,256 435,266 440,275l210,425l100,0l0,-700l-100,0l0,440C650,451 651,469 652,494l2,31l-4,0l-9,-31C639,489 636,480 632,469C628,457 624,447 620,440l-160,-315l-80,0l-160,315C216,447 212,457 208,469C204,480 201,489 199,494l-9,31l-4,0l2,-31C189,469 190,451 190,440l0,-440l-100,0z"/>
        <glyph unicode="N" horiz-adv-x="630" d="M90,700l110,0l210,-465C413,230 415,225 417,220C418,215 420,210 421,206l9,-25C432,175 434,169 436,164C437,159 439,154 440,150l4,0l-2,31C441,204 440,222 440,235l0,465l100,0l0,-700l-110,0l-210,465C214,478 207,496 200,519C199,523 197,528 196,533C194,538 192,544 190,550l-4,0l2,-31C189,494 190,476 190,465l0,-465l-100,0z"/>
        <glyph unicode="O" horiz-adv-x="600" d="M55,350C55,475 77,566 120,624C163,681 223,710 300,710C377,710 437,681 480,624C523,566 545,475 545,350C545,225 523,134 480,77C437,19 377,-10 300,-10C223,-10 163,19 120,77C77,134 55,225 55,350M445,350C445,452 431,526 404,572C377,617 342,640 300,640C258,640 223,617 196,572C169,526 155,452 155,350C155,248 169,174 196,129C223,83 258,60 300,60C342,60 377,83 404,129C431,174 445,248 445,350z"/>
        <glyph unicode="P" horiz-adv-x="565" d="M90,700l180,0C355,700 418,681 459,642C500,603 520,546 520,470C520,394 500,337 459,298C418,259 355,240 270,240l-80,0l0,-240l-100,0M270,310C321,310 358,323 383,349C408,375 420,415 420,470C420,525 408,565 383,591C358,617 321,630 270,630l-80,0l0,-320z"/>
        <glyph unicode="R" horiz-adv-x="575" d="M90,700l170,0C344,700 407,681 448,644C489,607 510,552 510,480C510,373 464,305 373,275l167,-275l-115,0l-155,260l-80,0l0,-260l-100,0M260,330C360,330 410,380 410,480C410,580 360,630 260,630l-70,0l0,-300z"/>
        <glyph unicode="S" horiz-adv-x="505" d="M89,17C67,26 49,35 35,45l35,60l10,0C89,98 103,91 121,82C157,67 193,60 230,60C275,60 309,71 334,93C358,114 370,143 370,180C370,203 364,223 351,240C338,256 323,269 304,280C285,291 260,304 227,319C188,336 156,353 132,368C108,383 88,403 71,428C54,453 45,483 45,520C45,557 54,589 71,618C88,647 112,669 144,686C175,702 212,710 255,710C311,710 360,701 401,682C419,675 435,666 450,655l-35,-60l-10,0C396,602 384,610 369,617C336,632 301,640 265,640C231,640 203,629 180,608C157,586 145,560 145,530C145,507 151,487 164,471C177,454 192,441 211,430C230,419 255,406 288,391C327,374 359,358 383,343C407,328 428,308 445,283C462,258 470,227 470,190C470,129 450,80 410,44C370,8 313,-10 240,-10C181,-10 131,-1 89,17z"/>
        <glyph unicode="T" horiz-adv-x="480" d="M190,0l0,630l-180,0l0,70l460,0l0,-70l-180,0l0,-630z"/>
        <glyph unicode="a" horiz-adv-x="510" d="M45,145C45,192 62,229 97,256C131,282 185,295 260,295l80,0l0,25C340,359 331,389 313,410C294,430 268,440 235,440C211,440 185,432 157,417C143,409 132,402 125,395l-10,0l-35,60C93,465 108,474 125,482C161,501 201,510 245,510C283,510 317,502 346,486C375,469 398,447 415,418C432,389 440,356 440,320l0,-320l-70,0l-10,60l-5,0C346,47 334,35 320,25C291,2 254,-10 210,-10C158,-10 118,4 89,32C60,59 45,97 45,145M310,95C321,104 331,114 340,125l0,100l-80,0C222,225 193,218 174,203C155,188 145,169 145,145C145,119 153,99 168,86C183,72 205,65 235,65C260,65 285,75 310,95z"/>
        <glyph unicode="c" horiz-adv-x="455" d="M50,250C50,335 69,400 107,444C144,488 195,510 260,510C304,510 344,501 380,482C397,474 412,465 425,455l-35,-60l-10,0C371,402 360,409 347,417C319,432 293,440 270,440C233,440 203,425 182,395C161,365 150,317 150,250C150,183 161,134 184,105C206,75 238,60 280,60C305,60 330,67 357,82C371,90 382,98 390,105l10,0l35,-60C423,35 408,26 390,17C355,-1 315,-10 270,-10C201,-10 147,12 108,56C69,100 50,165 50,250z"/>
        <glyph unicode="g" horiz-adv-x="535" d="M125,-183C107,-174 92,-165 80,-155l35,60l10,0C132,-102 142,-109 157,-118C184,-133 210,-140 235,-140C278,-140 309,-130 330,-109C350,-88 360,-55 360,-10l0,30C350,14 339,9 327,5C302,-5 277,-10 250,-10C190,-10 142,12 105,57C68,101 50,165 50,250C50,306 58,354 75,393C91,432 113,461 140,481C167,500 197,510 230,510C271,510 307,498 340,475C356,462 368,451 375,440l5,0l10,60l70,0l0,-510C460,-71 441,-119 403,-155C364,-192 312,-210 245,-210C200,-210 160,-201 125,-183M333,70C345,73 354,77 360,80l0,295C351,386 341,396 330,405C305,425 280,435 255,435C224,435 199,420 180,390C160,360 150,313 150,250C150,184 161,136 182,106C203,75 230,60 265,60C288,60 310,63 333,70z"/>
        <glyph unicode="l" horiz-adv-x="250" d="M75,0l0,670l30,30l70,0l0,-700z"/>
        <glyph unicode="n" horiz-adv-x="545" d="M75,500l70,0l10,-60l5,0C167,451 179,462 195,475C228,498 264,510 305,510C355,510 396,494 428,463C459,431 475,390 475,340l0,-340l-100,0l0,340C375,368 366,391 349,409C331,426 308,435 280,435C255,435 230,425 205,405C194,396 184,386 175,375l0,-375l-100,0z"/>
        <glyph unicode="o" horiz-adv-x="530" d="M50,250C50,335 69,400 107,444C145,488 198,510 265,510C332,510 384,488 423,444C461,400 480,335 480,250C480,165 461,100 423,56C384,12 332,-10 265,-10C198,-10 145,12 107,56C69,100 50,165 50,250M380,250C380,316 370,364 349,395C328,425 300,440 265,440C230,440 203,425 182,395C161,364 150,316 150,250C150,184 161,136 182,106C203,75 230,60 265,60C300,60 328,75 349,106C370,136 380,184 380,250z"/>
        <glyph unicode="t" horiz-adv-x="400" d="M95,150l0,280l-75,0l0,60l130,130l45,0l0,-120l125,0l0,-70l-125,0l0,-280C195,121 202,98 215,83C228,68 247,60 270,60C290,60 310,67 329,80C338,86 346,93 355,100l10,0l35,-60C385,29 372,20 360,15C328,-2 295,-10 260,-10C210,-10 170,5 140,34C110,63 95,102 95,150z"/>
        <glyph unicode="3" horiz-adv-x="500" d="M64,17C43,26 26,35 15,45l35,60l10,0C69,98 81,90 96,82C128,67 163,60 200,60C255,60 295,72 321,96C347,119 360,154 360,200C360,246 347,281 321,305C295,328 255,340 200,340l-105,0l0,70l225,220l-290,0l0,70l410,0l0,-70l-225,-220C296,410 357,392 398,355C439,318 460,266 460,200C460,133 439,81 398,45C357,8 294,-10 210,-10C151,-10 103,-1 64,17z"/>
        <glyph unicode="(" horiz-adv-x="285" d="M55,280C55,367 71,451 104,531C137,610 184,678 245,735l40,-40C242,634 210,570 188,503C166,436 155,362 155,280C155,198 166,124 188,57C210,-10 242,-74 285,-135l-40,-40C184,-118 137,-50 104,30C71,109 55,193 55,280z"/>
        <glyph unicode=")" horiz-adv-x="285" d="M0,-135C43,-74 75,-10 97,57C119,124 130,198 130,280C130,362 119,436 97,503C75,570 43,634 0,695l40,40C101,678 148,610 181,531C214,451 230,367 230,280C230,193 214,109 181,30C148,-50 101,-118 40,-175z"/>
    </font>

	<g id="BG" display="none">
	<g id="Polar_Grid_Graph" display="inline">
		<g id="POLAR">

				<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="388.7" ry="383.5"/>

				<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="323.9" ry="319.6"/>

				<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="259.1" ry="255.7"/>

				<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="194.3" ry="191.8"/>

				<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="129.6" ry="127.8"/>
			<ellipse fill="none" stroke="#666666" stroke-width="0.25" stroke-miterlimit="10" cx="391.6" cy="415.5" rx="64.8" ry="63.9"/>
		</g>
		<g id="BOLD">
			<line id="_x34_" fill="none" stroke="#666666" stroke-miterlimit="10" x1="391.6" y1="415.5" x2="585.4" y2="83.5"/>
			<line id="_x33_" fill="none" stroke="#666666" stroke-miterlimit="10" x1="391.6" y1="415.5" x2="690.4" y2="661.5"/>
			<line id="_x32_" fill="none" stroke="#666666" stroke-miterlimit="10" x1="391.6" y1="415.5" x2="140.4" y2="708.5"/>
			<line id="_x31_" fill="none" stroke="#666666" stroke-miterlimit="10" x1="391.6" y1="415.5" x2="147.4" y2="117.5"/>
		</g>
		<g id="BLUE">
			<polyline fill="none" stroke="#0000FF" stroke-width="0.5" stroke-miterlimit="10" points="210.4,76.1 391.4,415.5 449.5,36.2
							"/>
			<polyline fill="none" stroke="#0000FF" stroke-width="0.5" stroke-miterlimit="10" points="521.4,53.5 391.4,415.5 289.4,45.5
							"/>
		</g>
		<g id="RED">
			<polyline fill="none" stroke="#C1272D" stroke-width="0.5" stroke-miterlimit="10" points="3.4,400.5 391.4,415.5 87.4,653.5
				"/>
			<polyline fill="none" stroke="#C1272D" stroke-width="0.5" stroke-miterlimit="10" points="101.4,159.5 391.4,415.5 28.4,279.5
							"/>
			<line fill="none" stroke="#C1272D" stroke-width="0.5" stroke-miterlimit="10" x1="391.4" y1="415.5" x2="22.4" y2="534.5"/>
		</g>
		<g id="GREEN">
			<line fill="none" stroke="#39B54A" stroke-width="0.5" stroke-miterlimit="10" x1="202.4" y1="750.5" x2="391.4" y2="415.5"/>
			<polyline fill="none" stroke="#39B54A" stroke-width="0.5" stroke-miterlimit="10" points="499.4,784.5 391.4,415.5 343.4,796.5
							"/>
			<line fill="none" stroke="#39B54A" stroke-width="0.5" stroke-miterlimit="10" x1="391.4" y1="415.5" x2="631.4" y2="717.5"/>
		</g>
		<g id="YELLOW">
			<polyline fill="none" stroke="#F7931E" stroke-width="0.5" stroke-miterlimit="10" points="763.1,302.4 391.4,415.5 733.6,598
							"/>
			<polyline fill="none" stroke="#F7931E" stroke-width="0.5" stroke-miterlimit="10" points="780.2,458.5 391.4,415.5 706.1,190
							"/>
			<line fill="none" stroke="#F7931E" stroke-width="0.5" stroke-miterlimit="10" x1="636.4" y1="117.5" x2="391.4" y2="415.5"/>
		</g>
	</g>
</g>
<g id="LOGO" display="none">
</g>
<g id="Object">
	<g id="Green">
		<g id="MSLTP">
			<polygon fill="#9ACACE" points="271.9,636 418.9,733 414.9,740 265.9,644 			"/>
			<polygon fill="#9ACACE" points="454.9,600 418.9,733 412.9,733 441.9,598 			"/>
			<polygon fill="#9ACACE" points="379.9,550 424.9,743 414.9,740 370.9,550 			"/>
			<polygon fill="#9ACACE" points="520.9,575 428.9,736 418.9,733 510.9,572 			"/>
            <polygon fill="#9ACACE" points="299.9,584 424.9,736 417.9,740 293.9,592 			"/>
			<circle id="MSLTP_circle" fill="#9ACACE" cx="419.4" cy="735.5" r="22.5"/>
			<text id="course_34_" transform="matrix(0.9999 -1.566269e-02 1.566269e-02 0.9999 404.8404 733.2244)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MSL</tspan><tspan x="6.3" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">TP</tspan></text>
		</g>
		<g id="Ccl2">
			<circle id="Ccl2_circle" fill="#9ACACE" cx="443.4" cy="600.5" r="22.5"/>
			<text id="course_30_" transform="matrix(0.9999 -1.566269e-02 1.566269e-02 0.9999 433.4531 598.1522)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Ccl</tspan><tspan x="0.3" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(II)</tspan></text>
		</g>
		<g id="Ccl1">
			<polygon id="edge_26_" fill="#9ACACE" points="378.7,544.7 384.9,474 378.9,472 370.9,543.6 			"/>
			<circle id="Ccl1_circle" fill="#9ACACE" cx="374.4" cy="549.5" r="22.5"/>
			<text id="course_29_" transform="matrix(0.9999 -1.566269e-02 1.566269e-02 0.9999 364.4531 547.1522)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Ccl</tspan><tspan x="2.8" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="Stat2">
			<circle id="Stat2_circle" fill="#9ACACE" cx="265.4" cy="638.5" r="22.5"/>
			<text id="course_31_" transform="matrix(0.9999 -1.566269e-02 1.566269e-02 0.9999 250.6158 636.228)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Stat</tspan><tspan x="5.1" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(II)</tspan></text>
		</g>
		<g id="Stat1">
			<polygon id="edge_25_" fill="#9ACACE" points="302.3,581.1 362.9,471 358.9,468 295.5,576.8 			"/>
			<circle id="Stat1_circle" fill="#9ACACE" cx="296.3" cy="583.6" r="22.5"/>
			<text id="course_28_" transform="matrix(0.9996 2.987289e-02 -2.987289e-02 0.9996 281.732 581.0581)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Stat </tspan><tspan x="7.6" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="MM">
			<polygon id="edge_27_" fill="#9ACACE" points="513.1,563.5 432.9,466.8 428.6,469.4 506.7,568.2 			"/>
			<circle id="MM_circle" fill="#9ACACE" cx="513.4" cy="570" r="22.5"/>
	        <text id="course_32_" transform="matrix(0.9999 1.600062e-02 -1.600062e-02 0.9999 500.1757 575.1193)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MM</text>
		</g>
		<g id="DM">
			<polygon id="edge_28_" fill="#9ACACE" points="434.2,536.5 409.4,469.9 403.1,470.7 426.6,538.9 			"/>
			<circle id="DM_circle" fill="#9ACACE" cx="432.3" cy="542.6" r="22.5"/>
            <text id="course_33_" transform="matrix(0.9991 -4.275507e-02 4.275507e-02 0.9991 421.3832 548.1674)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">DM</text>
		</g>
	</g>
	<g id="Blue">
		<g id="IMFP">
			<polygon id="edge" fill="#99ACBB" points="310.9,247 362.8,353.9 355.5,357.9 301.9,252 			"/>
			<circle id="IMFP_circle" fill="#99ACBB" cx="305.6" cy="243.2" r="22.5"/>
            <text id="course_10_" transform="matrix(0.9678 -0.2516 0.2516 0.9678 289.207 253.4719)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">IMFP</text>
		</g>
        <g id="PoIM2">
			<polygon id="edge_12_" fill="#99ACBB" points="312.9,117 324.9,153 316.9,155 302.6,118.3 			"/>
			<circle id="PoIM2_circle" fill="#99ACBB" cx="306.2" cy="109.5" r="22.5"/>
			<text id="course_12_" transform="matrix(0.9678 -0.2516 0.2516 0.9678 288.6421 111.8712)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">PoIM</tspan><tspan x="8.1" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(II)</tspan></text>
		</g>
		<g id="PoIM1">
			<polygon id="edge_11_" fill="#99ACBB" points="332.9,177 378.9,353 370.9,354 322.9,181 			"/>
			<circle id="PoIM1_circle" fill="#99ACBB" cx="325.2" cy="167.5" r="22.5"/>
			<text id="course_11_" transform="matrix(0.9678 -0.2516 0.2516 0.9678 307.6568 169.8727)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">PoIM</tspan><tspan x="10.6" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="MPD1">
			<polygon id="edge_13_" fill="#99ACBB" points="424.9,227 404.9,350 396.9,348 415.9,225 			"/>
			<circle id="MPD1_circle" fill="#99ACBB" cx="420" cy="226.5" r="22.5"/>
			<text id="course_13_" transform="matrix(0.9994 3.569945e-02 -3.569945e-02 0.9994 404.1844 225.0512)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MPD </tspan><tspan x="8.5" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="MPD2">
			<polygon id="edge_14_" fill="#99ACBB" points="486.9,178 414.9,363 408.9,362 475.9,174 			"/>
			<circle id="MPD2_circle" fill="#99ACBB" cx="483" cy="175.5" r="22.5"/>
			<text id="course_14_" transform="matrix(0.9994 3.569945e-02 -3.569945e-02 0.9994 468.1844 173.0512)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MPD </tspan><tspan x="6.1" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(II)</tspan></text>
		</g>
		<g id="SILP">
			<circle id="SILP_circle" fill="#99ACBB" cx="383" cy="95.5" r="22.5"/>
            <text id="course_15_" transform="matrix(0.9994 3.569945e-02 -3.569945e-02 0.9994 367.9307 101.0549)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">SILP</text>
		</g>
	</g>
	<g id="Red">
		<g id="CN">
			<polygon id="edge_10_" fill="#FF9999" points="296.6,495.1 340.4,459.3 335.4,453.1 292.5,488.3 			"/>
			<circle id="CN_circle" fill="#FF9999" cx="290.2" cy="494.8" r="22.5"/>

				<text id="course_9_" transform="matrix(1 7.428121e-03 -7.428121e-03 1 280.531 500.5015)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">CN</text>
		</g>
		<g id="ALG">
			<polygon id="edge_9_" fill="#FF9999" points="148.8,498.2 202.5,480.2 199.9,472.6 147.4,490.4 			"/>
			<circle id="ALG_circle" fill="#FF9999" cx="142.9" cy="495.6" r="22.5"/>
            <text id="course_8_" transform="matrix(0.949 -0.3152 0.3152 0.949 133.2205 504.9724)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">ALG</text>
		</g>
		<g id="DS">
			<polygon id="edge_8_" fill="#FF9999" points="211.8,476.2 265.5,458.2 262.9,450.6 210.4,468.4 			"/>
			<circle id="DS_circle" fill="#FF9999" cx="205.9" cy="473.6" r="22.5"/>
	        <text id="course_7_" transform="matrix(0.949 -0.3152 0.3152 0.949 200.375 481.591)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">DS</text>
		</g>
		<g id="PD">
			<polygon id="edge_7_" fill="#FF9999" points="272.9,458.5 326.6,440.5 324,433 271.5,450.7 			"/>
			<circle id="PD_circle" fill="#FF9999" cx="267" cy="455.9" r="22.5"/>
            <text id="course_6_" transform="matrix(0.949 -0.3152 0.3152 0.949 261.043 464.1138)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">PD</text>
		</g>
		<g id="DB">
			<polygon id="edge_6_" fill="#FF9999" points="201.5,412.2 258,414.2 258.3,406.3 202.9,404.4 			"/>
			<circle id="DB_circle" fill="#FF9999" cx="196.9" cy="407.7" r="22.5"/>
            <text id="course_5_" transform="matrix(0.9992 3.967613e-02 -3.967613e-02 0.9992 188.3359 413.2012)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">DB</text>
		</g>
		<g id="OS">
			<polygon id="edge_5_" fill="#FF9999" points="266.3,415.9 322.9,416.8 323.1,408.9 267.6,408.1 			"/>
			<circle id="OS_circle" fill="#FF9999" cx="261.7" cy="411.5" r="22.5"/>
            <text id="course_4_" transform="matrix(0.9998 1.919349e-02 -1.919349e-02 0.9998 253.6423 417.2397)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">OS</text>
		</g>
		<g id="COS">
			<polygon id="edge_4_" fill="#FF9999" points="273.2,373.4 325.8,394.3 328.7,387 277.1,366.5 			"/>
			<circle id="COS_circle" fill="#FF9999" cx="270.4" cy="367.6" r="22.5"/>
            <text id="course_3_" transform="matrix(1 0 0 1 257.7298 374.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">COS</text>
		</g>
		<g id="SPM">
			<circle id="SPM_circle" fill="#FF9999" cx="196.4" cy="243.5" r="22.5"/>
            <text id="course_2_" transform="matrix(1 0 0 1 181.554 249.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">SPM</text>
		</g>
		<g id="SDM">
			<polygon id="edge_2_" fill="#FF9999" points="247.2,292.9 290.5,329.4 295.6,323.3 253.2,287.6 			"/>
			<circle id="SDM_circle" fill="#FF9999" cx="246.4" cy="286.5" r="22.5"/>
            <text id="course_1_" transform="matrix(1 0 0 1 231.4046 292.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">SDM</text>
		</g>
		<g id="SA">
			<polygon id="edge_1_" fill="#FF9999" points="295.2,335.9 338.5,372.4 343.6,366.3 301.2,330.6 			"/>
			<circle id="SA_circle" fill="#FF9999" cx="294.4" cy="329.5" r="22.5"/>
            <text id="course" transform="matrix(1 0 0 1 286.1917 335.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">SA</text>
		</g>
		<g id="IR">
			<circle id="IR_circle" fill="#FF9999" cx="155.4" cy="321.5" r="22.5"/>
            <text id="course_35_" transform="matrix(1 0 0 1 148.8421 327.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">IR</text>
		</g>
		<g id="GP">
			<circle id="GP_circle" fill="#FF9999" cx="124.4" cy="235.5" r="22.5"/>
            <text id="course_36_" transform="matrix(1 0 0 1 115.6292 241.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">GP</text>
		</g>
		<g id="3DP">
			<circle id="3DP_circle" fill="#FF9999" cx="90.4" cy="297.5" r="22.5"/>
            <text id="course_37_" transform="matrix(1 0 0 1 77.5042 303.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">3DP</text>
		</g>
		<g id="BDBA">
			<circle id="BDBA_circle" fill="#FF9999" cx="74.4" cy="367.5" r="22.5"/>
            <text id="course_38_" transform="matrix(1 0 0 1 56.2171 373.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">BDBA</text>
		</g>
		<g id="3DCR">
			<circle id="3DCR_circle" fill="#FF9999" cx="67.4" cy="445.5" r="22.5"/>
            <text id="course_39_" transform="matrix(1 0 0 1 50.1917 451.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">3DCR</text>
		</g>
		<g id="IS">
			<circle id="IS_circle" fill="#FF9999" cx="183.4" cy="574.5" r="22.5"/>
            <text id="course_40_" transform="matrix(1 0 0 1 177.3665 580.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">IS</text>
		</g>
		<g id="ToC">
			<circle id="ToC_circle" fill="#FF9999" cx="82.4" cy="517.5" r="22.5"/>
            <text id="course_41_" transform="matrix(1 0 0 1 70.4417 523.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">ToC</text>
		</g>
		<g id="PL">
			<circle id="PL_circle" fill="#FF9999" cx="238.4" cy="535.5" r="22.5"/>
            <text id="course_42_" transform="matrix(1 0 0 1 230.3421 541.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">PL</text>
		</g>
		<g id="IE">
			<circle id="IE_circle" fill="#FF9999" cx="104.4" cy="572.5" r="22.5"/>
	        <text id="course_43_" transform="matrix(1 0 0 1 98.3294 578.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">IE</text>
		</g>
		<g id="HCI">
			<circle id="HCI_circle" fill="#FF9999" cx="140.4" cy="619.5" r="22.5"/>
            <text id="course_44_" transform="matrix(1 0 0 1 129.1917 625.3071)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">HCI</text>
		</g>
	</g>
	<g id="Yellow">
		<g id="FM">
			<polygon id="edge_16_" fill="#F8DA9E" points="570.7,241.1 517.5,260.3 520.2,267.8 572.3,248.9 			"/>
            <polygon id="edge_17_" fill="#F8DA9E" points="569.6,243.7 555.6,298.5 563.3,300.5 577,246.7 			"/>
			<circle id="FM_circle" fill="#F8DA9E" cx="579.4" cy="238.5" r="22.5"/>
            <text id="course_19_" transform="matrix(0.9859 0.1675 -0.1675 0.9859 566.5363 241.5181)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">FM</text>
		</g>
		<g id="MIS">
			<polygon id="edge_3_" fill="#F8DA9E" points="468.3,317.4 431.3,360.2 437.3,365.4 473.5,323.4 			"/>
			<circle id="MIS_circle" fill="#F8DA9E" cx="474.7" cy="316.6" r="22.5"/>
            <text id="course_16_" transform="matrix(0.9859 0.1675 -0.1675 0.9859 459.3513 319.2076)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MIS</text>
		</g>
		<g id="Acct1">
			<polygon id="edge_15_" fill="#F8DA9E" points="543,303.6 445.9,373 448.9,377 546.4,310.8 			"/>
			<circle id="Acct1_circle" fill="#F8DA9E" cx="549.4" cy="304.6" r="22.5"/>
			<text id="course_17_" transform="matrix(0.9999 -1.279673e-02 1.279673e-02 0.9999 533.2446 303.2374)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Acct </tspan><tspan x="8.1" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="Acct2">
			<circle id="Acct2_circle" fill="#F8DA9E" cx="513.4" cy="267.6" r="22.5"/>
			<text id="course_18_" transform="matrix(0.9999 -1.279673e-02 1.279673e-02 0.9999 497.2484 266.2226)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Acct </tspan><tspan x="5.6" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(II)</tspan></text>
		</g>
		<g id="Eco1">
			<polygon id="edge_18_" fill="#F8DA9E" points="512,373.8 458,390.7 460.4,398.3 513.3,381.7 			"/>
			<circle id="Eco1_circle" fill="#F8DA9E" cx="517.9" cy="376.5" r="22.5"/>
			<text id="course_20_" transform="matrix(0.9987 5.127607e-02 -5.127607e-02 0.9987 506.2646 373.671)"><tspan x="0" y="0" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Eco </tspan><tspan x="4.8" y="18" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">(I)</tspan></text>
		</g>
		<g id="SM">
			<polygon id="edge_23_" fill="#F8DA9E" points="639,387.7 520.7,422.2 522.5,429 640.2,395.6 			"/>
			<circle id="SM_circle" fill="#F8DA9E" cx="644.8" cy="390.5" r="22.5"/>
            <text id="course_25_" transform="matrix(0.9989 -4.760430e-02 4.760430e-02 0.9989 635.0378 396.568)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">SM</text>
		</g>
		<g id="MtM">
			<polygon id="edge_24_" fill="#F8DA9E" points="632.1,490.5 524.7,430.1 521,436.2 627.3,496.8 			"/>
			<circle id="MtM_circle" fill="#F8DA9E" cx="634.1" cy="496.6" r="22.5"/>
	        <text id="course_26_" transform="matrix(1 -2.194629e-05 2.194629e-05 1 617.1284 502.3309)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">MtM</text>
		</g>
		<g id="OB">
			<polygon id="edge_22_" fill="#F8DA9E" points="643.5,440.7 520.9,428 519.9,435 641.6,448.4 			"/>
			<circle id="OB_circle" fill="#F8DA9E" cx="647.8" cy="445.5" r="22.5"/>
            <text id="course_24_" transform="matrix(0.9998 -1.747232e-02 1.747232e-02 0.9998 638.5879 451.4417)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">OB</text>
		</g>
		<g id="Mng">
			<polygon id="edge_19_" fill="#F8DA9E" points="514.5,422.7 458.2,417 457.4,424.9 512.6,430.4 			"/>
			<circle id="Mng_circle" fill="#F8DA9E" cx="518.8" cy="427.5" r="22.5"/>
            <text id="course_21_" transform="matrix(0.9998 -1.747232e-02 1.747232e-02 0.9998 503.6313 433.5442)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">Mng</text>
		</g>
		<g id="OM">
			<polygon id="edge_21_" fill="#F8DA9E" points="564.9,501.7 513.4,478 510.1,485.2 560.6,508.3 			"/>
			<circle id="OM_circle" fill="#F8DA9E" cx="567.4" cy="507.6" r="22.5"/>
            <text id="course_23_" transform="matrix(0.9981 -6.208936e-02 6.208936e-02 0.9981 555.9271 514.0656)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">OM</text>
		</g>
		<g id="OR">
			<polygon id="edge_20_" fill="#F8DA9E" points="503.9,469.7 452.5,446 449.1,453.2 499.6,476.4 			"/>
			<circle id="OR_circle" fill="#F8DA9E" cx="506.4" cy="475.6" r="22.5"/>
            <text id="course_22_" transform="matrix(0.9981 -6.208936e-02 6.208936e-02 0.9981 496.9164 481.9635)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">OR</text>
		</g>
		<g id="ISIM">
			<circle id="ISIM_circle" fill="#F8DA9E" cx="705.4" cy="315.5" r="22.5"/>
            <text id="course_27_" transform="matrix(0.9999 -1.242842e-02 1.242842e-02 0.9999 690.1734 321.5838)" fill="#FFFFFF" font-family="'Scada-Regular'" font-size="15px">ISIM</text>
		</g>
	</g>
	<g id="Center">
		<path fill="#CC3333" stroke="#FFFFFF" stroke-width="5" stroke-miterlimit="10" d="M341.1,423.7c-3-18.6,4.6-36.5,18.5-47.4
			l-11-13.4c-3.2,2.6-5.7,5.1-8.4,8.2c-24.4,28.3-21.2,71,7.1,95.3l10.9-12.6C349.4,446.4,343.1,435.9,341.1,423.7z"/>
		<path fill="#008490" stroke="#FFFFFF" stroke-width="5" stroke-miterlimit="10" d="M398.5,465c-15,2.5-29.6-2.1-40.4-11.2
			l-10.9,12.6c3.1,2.7,6,4.8,9.6,6.9c32.1,19.1,73.6,8.5,92.6-23.6l-15.6-9.3C426.7,453.1,414,462.5,398.5,465z"/>
		<path fill="#FF9933" stroke="#FFFFFF" stroke-width="5" stroke-miterlimit="10" d="M415.4,352l-6.5,17.2c15.8,6.3,28,20.4,31,38.5
			c1.9,11.8-0.4,23.3-5.9,32.9l15.6,9.3c2.1-3.6,3.6-6.7,5.1-10.6C467.9,404.2,450.3,365.2,415.4,352z"/>
		<path fill="#006699" stroke="#FFFFFF" stroke-width="5" stroke-miterlimit="10" d="M382.4,366.3c9.2-1.5,18.3-0.4,26.4,2.8
			l6.5-17.2c-23.3-8.8-47.5-4.9-66.7,10.9l11,13.4C366.1,371.3,373.8,367.7,382.4,366.3z"/>
	</g>
</g>
</svg>
`

var tree02state = `
{	
	"PD":{"CFullName":"程式設計",CName":"程設","EFullName":"Programming Design","EName":"PD","pre_course":[],"status":0,"color":"red","CDescription":"本課程涵蓋計算機程式編寫之入門知識。本課程選用程式語言為C++，首先由C++基礎程序起步，接著討論C++的物件導向特性。除了訓練如何編寫「正確的」程式，本課程也著重於編寫「良好的」程式，意即較為快速、運用較少記憶體、使用者介面友善、易於延伸、具有較好格式之程式。此外，為提高學習效率，本課程也將討論資料結構、計算複雜性及演算法設計的基礎觀念。","EDescription":"In this course, we will introduce how to write computer programs for general purposes. The programming language we will study is C++. We will start from the procedural programming part of C++, and then discuss those object-oriented features of C++. While we will spend a lot of time on how to write correct programs, we will also aim to write “good” programs, i.e., those running faster, using less memory, generating friendly user interfaces, being more extendable, having better formats, etc. To enhance the learning efficiency, basic concepts of data structure, computational complexity, and algorithm design will also be discussed."},
	"DS":{"CFullName":"資料結構與進階程式設計","CName":"資結","EFullName":"Data Structures and Advanced Program Design","EName":"DS","pre_course":["PD"],"status":0,"color":"red","CDescription":"本課程介紹抽象資料形式的基本觀念，使學生熟知個常見抽象資料形式的設計原理及程式應用，例如：堆疊、優先佇列、二元搜尋樹等，並透過仔細的討論，了解這些資料結構以及其背後操作之原理。","EDescription":"This course introduces to the students the fundamental concept of an abstract data type. Its goal is to acquaint the students with the design principles and applications of several commonly used abstract data types such as stacks, priority queues, and binary search trees and, through close examination, the data structures and their manipulations that underly the implementations of these abstract data types."},
	"PL":{"CFullName":"程式語言","CName":"程式語言","EFullName":"Programming Languages ","EName":"PL","pre_course":[],"status":0,"color":"red","CDescription":"本課程將使學生瞭解程式語言在語法、語意、與實務面向的課題，作為日後輔助思維與規劃的依據。","EDescription":"This course is designed to familiarize students with issues related to programming language in syntax, semantic and practical fields, providing assistance in thinking and planning."},
	"ToC":{"CFullName":"計算理論","CName":"計理","EFullName":"Theory of Computing","EName":"ToC","pre_course":[],"status":0,"color":"red","CDescription":"本課程著重於使學生熟悉計算理論的基礎觀念，並培養學生分析問題複雜度的能力。","EDescription":"The goal of this course is to acquaint the students with the basic concepts in computation theory and to cultivate the students' ability in analyzing the complexity of computational problems."},
	"CN":{"CFullName":"網路技術與應用","CName":"網路","EFullName":"Computer Networks and Applications","EName":"CN","pre_course":[],"status":0,"color":"red","CDescription":"本課程將介紹電腦網路如何運作、發展，尤其著重於網際網路、網路服務以及網路應用，將由網路分層層層遞進，介紹各層所代表的意義與運作原理。此外，注重培養學生察覺問題的能力，並且釐清各種解決方法的邏輯意義。","EDescription":"This introductory course is intended to give students general knowledge on how computer networks work, especially on the Internet and the field of networking services and applications. Special focus is on motivating students how the problems were raised and the rationale behind various solution approaches."},
	"IS":{"CFullName":"資訊安全","CName":"資安","EFullName":"Information Security","EName":"IS","pre_course":[],"status":0,"color":"red","CDescription":"本課程涵蓋多使用者資訊系統與電腦網路中的資訊安全問題，將介紹基礎資訊安全技術，並討論實作上的應用，如企業內部網路、電子商務中的資安問題。","EDescription":"This course concerns security issues in multi-user information systems and computer networks. It will cover fundamental techniques for security and their applications in practical areas such as intranets and electronic commerce."},
	"OS":{"CFullName":"作業系統","CName":"作業系統","EFullName":"Operating Systems","EName":"OS","pre_course":[],"status":0,"color":"red","CDescription":"本課程涵蓋作業系統背後的概念及演算法，且為了讓學生更透徹了解現行系統的運作方式，課程中將介紹現今主流之作業系統，如微軟系列作業系統及各代UNIX/LINUX作業系統。","EDescription":"The goal of the course is to present the concepts and algorithms that underlie operating systems. In addition, to help students better understand the operation of modern systems, the concepts and algorithms covered in the course are often based on those used in existing commercial operating systems. Particular attention is paid to the microsoft family of operating systems and various versions of UNIX/LINUX."},
	"DB":{"CFullName":"資料庫管理","CName":"資料庫","EFullName":"Database Management","EName":"DB","pre_course":["OS"],"status":0,"color":"red","CDescription":"本課程將介紹資料庫建置及系統管理的基本觀念與原則。","EDescription":"This course aims at developing an understanding of basic concepts and principles of database management systems."},
	"COS":{"CFullName":"計算機組織與結構","CName":"計組","EFullName":"Computer Organization and Structure","EName":"COS","pre_course":[],"status":0,"color":"red","CDescription":"本課程涵蓋計算機組織與結構的數種基礎觀念，主題包含：記憶體與處理器之組織、指令集架構、組合語言、技術系統、資料傳輸與控制流指令、程序、記憶體管理、程式執行等。","EDescription":"This course intended to introduce basic concepts of computer organization and structure, including basic memory and processor organization, instruction set architecture, assembly language programming, number systems, arithmetic, data transfer and control flow instructions, procedures, memory management and program execution."},
	"SA":{"CFullName":"系統分析與設計","CName":"系統分析","EFullName":"Systems Analysis and Design","EName":"SA","pre_course":[],"status":0,"color":"red","CDescription":"本課程涵蓋電腦資訊系統開發之原則與方法，將討論各開發階段，如：計畫、分析、設計、測試階段之中的各式重要議題，並介紹各式重要系統分析與設計工具，如需求調查、需求定義、流程圖、詞彙字典及物件導向系統分析。","EDescription":"This course is designed to provide the overview of the principles and methodologies for the development of computer information systems. Important issues in different system development stages, including planning, analysis, design, testing, implication, and maintenance will be discussed. Use of various software engineering analysis and design tools and techniques are covered, including information gathering for defining system requirements, data flow diagrams, data dictionaries, and object-oriented systems analysis."},
	"SDM":{"CFullName":"軟體開發方法","CName":"軟開","EFullName":"Software Development Methods","EName":"SDM","pre_course":["SA"],"status":0,"color":"red","CDescription":"本課程將介紹一系列與開發軟體相關之理論與實作，使學生熟悉在軟體開發與軟體驗證領域廣泛採用之方法與工具，為未來職場所需之軟體開發技能打下堅實基礎。","EDescription":"This course introduces a selection of theories and practices that can enhance the student's ability in developing correct and high-quality software. Its goal is to acquaint the students with the well-used methods and tools for practical software development as well as some fundamentals of software verification, so as to prepare them for a career in software development."},
	"SPM":{"CFullName":"軟體專案管理","CName":"軟專","EFullName":"Software Project Management","EName":"SPM","pre_course":[],"status":0,"color":"red","CDescription":"本課程將透過三階段主題來使學生熟悉軟體開發：軟體專案計畫、軟體專案部署管理、軟體專案結案。","EDescription":"This course will train students to be familar with project management disciplines through three phases: software project planning, control of software project deployment, software project closure. Upon the completion of this course, students should be able to understand the basic knowledge and technical background of software development and IT project management."},
	"IR":{"CFullName":"資訊檢索與文字探勘導論","CName":"文字探勘","EFullName":"Introduction to Information Retrieval and Text Mining","EName":"IR","pre_course":[],"status":0,"color":"red","CDescription":"本課程主要面向所有對資料檢索與文字探勘有興趣的大學部學生及研究所學生。首先，本課程將介紹資料檢索的基礎觀念，主題包含：文件表示、字彙權重，以及用於比較資料探勘系統的評估標準。接著，將討論文字分類與群聚，對資料檢索與文字探勘進行更全面的學習。","EDescription":"The course is aimed at graduate students or senior undergraduate students who are interested in information retrieval and text mining. The first part of the course will cover the basics of information retrieval, including the manner of document representation, term weighting, and evaluation metrics used to compare information retrieval systems. Then, research topics, such as text classification and clustering, will be discussed to provide a comprehensive study on information retrieval and text mining."},
	"BDBA":{"CFullName":"大數據與商業分析","CName":"大數據","EFullName":"Big Data and Business Analytics","EName":"BDBA","pre_course":[],"status":0,"color":"red","CDescription":"本課程將著重於大數據與商業分析之理論基礎與架構，使學生了解各種大數據與商業分析系統之內涵並學習如何建置之，此外，亦將討論數位時代下最新的大數據與商業分析議題。","EDescription":"This course focuses on the fundamental principles and frameworks of Big Data and Business Analytics, acquainting students with the underlying essence of Big Data and Business Analytics, and thus, learn the methodology to build one by themselves. In addition, the state-of-art issues about Big Data and Business Analytics in the nowadays digital era will also be covered."},
	"GP":{"CFullName":"遊戲設計","CName":"遊戲設計","EFullName":"Game Programming","EName":"GP","pre_course":[],"status":0,"color":"red","CDescription":"","EDescription":""},
	"3DCR":{"CFullName":"3D創作實境","CName":"3D實境","EFullName":"3D Creation In Reality","EName":"3DCR","pre_course":[],"status":0,"color":"red","CDescription":"本課程將介紹迷人的3D世界，並連結學院上與市場上的3D技術。課程主題包含：3D基本知識與技術、3D作品創作練習、3D個案研究、3D作品期末呈現。","EDescription":"This course will introduce the captivating 3D world and attempts to fill the gap of 3D talents between the academic and the market place. Discussing topics covers 4 parts: 3D foundational knowledge and skill, 3D creation practice, 3D application case study and 3D practical term project."},
	"3DP":{"CFullName":"3D程式設計","CName":"3D程設","EFullName":"Introduction to 3D Programming","EName":"3DP","pre_course":[],"status":0,"color":"red","CDescription":"本課程將介紹3D電腦圖像、動畫、遊戲以及3D虛擬實境(VR)、擴增實境(AR)，使學生具備3D基礎知識，並培養學生的3D程式設計技術，如：HTML5 Canvas、HTML5 SVG、WebGL、P5.js及其應用、Three.js 3D JavaScript library等。","EDescription":"This course aims at introducing 3D programming to students who are interested in 3D computer graphics, animations, game development, 3D in real life as well as in VR and AR applications. Students wiil be eqquiped with both 3D fundamental knowledge and programming skills such as HTML5 Canvas, HTML5 SVG, WebGL, P5.js and Processing, Three.js 3D JavaScript library."},
	"IE":{"CFullName":"資訊經濟","CName":"資經","EFullName":"Information Economy","EName":"IE","pre_course":[],"status":0,"color":"red","CDescription":"本課程將帶領學生運用經濟模型，對各種資訊相關議題進行嚴謹的分析，尤其著重於資訊不對稱。本課程討論議題涵蓋：行銷、供應鏈管理、資訊系統等不同階段之應用。學生將透過閱讀課本及論文來進行學習，並於本課程中學習如何運用一特定研究方法進行學術研究。","EDescription":"In this course, we will study how to apply economic modeling to rigorously analyze information-related issues, especially information asymmetry. Applications that we will study lie in marketing, supply chain management, information systems, among others. Students will be required to read textbooks as well as some academic papers. This is a course teaching students how to do academic research with a specific research method."},
	"HCI":{"CFullName":"人機互動","CName":"人機","EFullName":"Human-Computer Interaction","EName":"HCI","pre_course":[],"status":0,"color":"red","CDescription":"","EDescription":""},
	"ALG":{"CFullName":"演算法","CName":"演算法","EFullName":"Algorithms ","EName":"ALG","pre_course":["DS","PD"],"status":0,"color":"red","CDescription":"本課程為計算機演算法設計與分析之入門課程，特色在於強調數學歸納法原理及其於演算法設計之應用，經由本課程的訓練，使學生熟悉基本之演算法及其設計原理，並培養學生具備獨立設計與分析演算法的能力。","EDescription":"This is an introductory course of algorithm design and analysis, focusing on mathematical induction and its application on algorithm design. Through this course, students would be acquaint with basic algorithms and the rationale lying behind. Also, students would cultivate the ability to design and analyze algorithms by themselves."},
	"Ccl1":{"CFullName":"微積分甲上","CName":"微甲上","EFullName":"Calculus (general Mathematics) (a)(1)","EName":"Ccl1","pre_course":[],"status":0,"color":"green","CDescription":"本課程涵蓋極限、微分、單變數積分、多變數積分等重要數學概念，有助於眾多科學、工程領域中的應用，學生亦具備微分方程、線性代數、與進階微積分的基礎。","EDescription":"After completing this course, students should be well versed in the mathematical language needed for applying the concepts of calculus to numerous applications in science and engineering. They should also be well prepared for courses in differential equations, linear algebra, or advanced calculus."},
	"Ccl2":{"CFullName":"微積分甲下","CName":"微甲下","EFullName":"Calculus (general Mathematics) (a)(2)","EName":"Ccl2","pre_course":["Ccl1"],"status":0,"color":"green","CDescription":"本課程涵蓋極限、微分、單變數積分、多變數積分等重要數學概念，有助於眾多科學、工程領域中的應用，學生亦具備微分方程、線性代數、與進階微積分的基礎。","EDescription":"After completing this course, students should be well versed in the mathematical language needed for applying the concepts of calculus to numerous applications in science and engineering. They should also be well prepared for courses in differential equations, linear algebra, or advanced calculus."},
	"Stat1":{"CFullName":"統計學一上","CName":"統計上","EFullName":"Statistics (I)(1)","EName":"Stat1","pre_course":[],"status":0,"color":"green","CDescription":"本課程著重於介紹統計之應用及基本觀念，並提供實務教學指引，讓學生學習如何選擇正確統計方法、進行統計計算，以及依據問題情境正確地解讀統計結果。鑑於商業決策皆是由大量數據所驅動，修習完本課程的學生將有能力於各種商業領域中，透過統計工具的幫助，進行有效、洞見性的決策。","EDescription":"This course emphasizes applications and fundamental concepts of statistics as well as provides a practical orientation that teaches students how to identify the correct method, calculate the statistics, and properly interpret the results in the context of the question or decision at hand. Since business decisions are driven by data, students will leave this course equipped with the tools they need to make effective, informed decisions in all areas of the business world."},
	"Stat2":{"CFullName":"統計學一下","CName":"統計下","EFullName":"Statistics (Ⅰ)(2)","EName":"Stat2","pre_course":["Stat1"],"status":0,"color":"green","CDescription":"本課程著重於介紹統計之應用及基本觀念，並提供實務教學指引，讓學生學習如何選擇正確統計方法、進行統計計算，以及依據問題情境正確地解讀統計結果。鑑於商業決策皆是由大量數據所驅動，修習完本課程的學生將有能力於各種商業領域中，透過統計工具的幫助，進行有效、洞見性的決策。","EDescription":"This course emphasizes applications and fundamental concepts of statistics as well as provides a practical orientation that teaches students how to identify the correct method, calculate the statistics, and properly interpret the results in the context of the question or decision at hand. Since business decisions are driven by data, students will leave this course equipped with the tools they need to make effective, informed decisions in all areas of the business world."},
	"MM":{"CFullName":"管理數學","CName":"管數","EFullName":"Management Mathematics","EName":"MM","pre_course":[],"status":0,"color":"green","CDescription":"本課程著重與線性代數相關的理論、演算法和應用。課程的主題包括線性方程式、矩陣、線性轉換、線性相依、行列式、向量空間、內積空間、特徵值、特徵向量等。本課程將引進抽象數學的概念、提供學生線性代數的基礎，使學生熟悉理論推導並培養學生運用線性代數代解決相關問題的能力。","EDescription":"This course focuses on the theories, algorithms and applications related to linear algebra. Topics of this course including: linear equations, matrices, linear transformation, linear dependence, determinant, real vector spaces,  inner product spaces, eigenvalues and eigenvectors. This course introduces the concepts of abstract mathematics, providing students with foundations of linear algebra, familiarizing students with theoretical derivation and the ability to solve problems with linear algebra."},
	"DM":{"CFullName":"離散數學","CName":"離散","EFullName":"Discrete Mathematics","EName":"DM","pre_course":[],"status":0,"color":"green","CDescription":"本課程目標為訓練學生的邏輯性思考與數理性思考，主要透過以下五大領域進行：邏輯、集合論、數論、圖論及應用模型。","EDescription":"The goal of this course is to train students to think logically and mathematically with the focus on the following 5 themes: mathematical reasoning, combinatorial analysis, discrete structure, algorithmic thinking, and applications and modeling."},
	"MSLTP":{"CFullName":"統計學習初論","CName":"統初","EFullName":"Modern Statistical Learning：Theory and Practic","EName":"MSLTP","pre_course":["Ccl1","Ccl2","Stat1","Stat2","MM"],"status":0,"color":"green","CDescription":"本課程將介紹如何運用統計學習與機器學習工具進行資料分析，課程中將涵蓋各種資料分析的工具，其中，學生非僅僅使用工具而已，而將會學習到這些工具背後之意義與細節，使學生能徹底了解資料分析工具運作的原理，從而有能力於各式各樣的選擇中採用最適合的工具來進行資料分析。","EDescription":"This course introduces the methodologies of using statistical learning and machine learning tools to analyze data. The goal of this course is to introduce a set of tools for data analytics. In addition, students will be exposed to the details instead of simply using these tools, enabling students to fully understand how a tool work and thus select the best tool among various choices."},
	"Acct1":{"CFullName":"會計學甲一上","CName":"會計上","EFullName":"Financial Accounting (A)(I)(1)","EName":"Acct1","pre_course":[],"status":0,"color":"yellow","CDescription":"本課程涵蓋財務會計的原理與處理程序，包括會計與商業活動的關係、財務報表的種類及其目的、會計恆等式、借貸法則、會計循環、內部控制、商業活動對資產、負債與權益（及收益、費損）科目的影響，乃至財務報表分析。","EDescription":"This course is an introduction to business financial reporting designed to create an awareness of the accounting principles and concepts for preparing and understanding the four basic financial statements: the balance sheet, statement of comprehensive income, statement of changes in equity, and statement of cash flows."},
	"Acct2":{"CFullName":"會計學甲一下","CName":"會計下","EFullName":"Financial Accounting (A)(I)(2)","EName":"Acct2","pre_course":["Acct1"],"status":0,"color":"yellow","CDescription":"本課程涵蓋財務會計的原理與處理程序，包括會計與商業活動的關係、財務報表的種類及其目的、會計恆等式、借貸法則、會計循環、內部控制、商業活動對資產、負債與權益（及收益、費損）科目的影響，乃至財務報表分析。","EDescription":"This course is an introduction to business financial reporting designed to create an awareness of the accounting principles and concepts for preparing and understanding the four basic financial statements: the balance sheet, statement of comprehensive income, statement of changes in equity, and statement of cash flows."},
	"Eco1":{"CFullName":"經濟學一","CName":"經濟","EFullName":"Economics (Ⅰ)","EName":"Eco1","pre_course":[],"status":0,"color":"yellow","CDescription":"本課程涵蓋經濟學之核心概念及決策意涵，並討論消費者、企業及政府的決策制定，以及經濟個體之決策所產生的機會成本。本課程討論議題包含：需求供給、國際貿易、外部性與公共財、消費者行為、廠商行為、市場結構等。","EDescription":"This course will introduce the core concepts and policy implications of economics. We shall discuss the decision making of choices by consumers, business and government, and the opportunity costs of one decision-making by economic units. Topics of this course includ demand and supply, international trade, externality and public goods, consumer behavior, firm behavior, market structure, and factor markets etc. and other topics on Microeconomics."},
	"Mng":{"CFullName":"管理學","CName":"管理學","EFullName":"Management","EName":"Mng","pre_course":[],"status":0,"color":"yellow","CDescription":"本課程介紹管理策略與技術的基本觀念，使學生學習各式各樣的管理方法，並從中了解該管理方法對組職本身、全球市場帶來的影響。本課程涵蓋課題包括：管理觀念的演化、領導能力、人力資源管控及組織文化。","EDescription":"This course aims to introduce students the fundamental concepts of management strategies and techniques. This course assists students to learn various management practices and how such practices might impact on the operation of organizations in the global marketplace. Topics of the course include evolution of management thought, leadership, human resource management and organizational culture."},
	"OR":{"CFullName":"作業研究","CName":"作業研究","EFullName":"Operations Research","EName":"OR","pre_course":[],"status":0,"color":"yellow","CDescription":"","EDescription":""},
	"MIS":{"CFullName":"資訊管理導論","CName":"資管導","EFullName":"Introduction to Information Management","EName":"MIS","pre_course":[],"status":0,"color":"yellow","CDescription":"本課程將涵蓋有關資訊系統的全面性概要，其中包含技術面與管理面之議題，此外，本課程也將討論資訊系統的基本觀念及其如何影響企業公司，並了解現今的資訊科技架構與重要技術，如：資料庫管理、商業智慧、通信系統與資訊系統安全。","EDescription":"This course offers a comprehensive overview of information systems (IS). It will explore both technical and managerial issues related to IS. Fundamental concepts of IS and how they impact business firms will be discussed. Students will also be introduced to modern information technology infrastructure and important components such as database management, business intelligence, telecommunications, and information systems security."},
	"ISIM":{"CFullName":"資訊服務創新管理","CName":"資管創","EFullName":"Information Service Innovation Management","EName":"ISIM","pre_course":[],"status":0,"color":"yellow","CDescription":"本課程的重點在介紹利用資訊科技來提供服務之創新管理，並探討各種資訊服務創新的相關策略，涵蓋議題包括：服務經濟與資訊服務創新、資訊服務創新的產業動態、資訊服務之創新與開發、創新管理、資訊服務創新之策略規劃、資訊服務創新策略之執行。","EDescription":"This course focuses on the innovative managements related to the services applying information technologies and discuss the associated strategies. Topics covered in this course include service economy and information service innovation, the industrial dynamics in information service innovation, the innovation and development of information services, innovative management, the strategy planning of information service innovation, the implementation of information service innovation."},
	"FM":{"CFullName":"財務管理","CName":"財管","EFullName":"Financial Management","EName":"FM","pre_course":["Acct1","Acct2"],"status":0,"color":"yellow","CDescription":"本課程將討論公司財務相關之理論與實作，討論主題涵蓋：財務決策、基本評價、投資決策、風險/報酬關聯、資本結構、支付政策、長期融資等議題。","EDescription":"This course is designed to provide a framework for understanding the theories and practices of corporate finance. The contents covered in this course include financial decision making, basic valuation, investment decision rules, risk-return relation, capital structure, payout policy, long-term financing, and other special topics."},
	"SM":{"CFullName":"策略管理","CName":"策管","EFullName":"Strategic Management","EName":"SM","pre_course":["Mng"],"status":0,"color":"yellow","CDescription":"本課程涵蓋公司組織之策略分析、構想及實行，將介紹產業競爭分析及公司現況分析之架構與工具，使學生同時具備產業層面及公司層面的決策能力。","EDescription":"This course is designed to study the processes of strategy analysis, formulation, and implementation for business organizations. This course will provide frameworks and tools for industry-level competition analysis and firm-level situation analysis. Class participants will apply what they learn to strategic decision making at both business level and corporate level."},
	"OB":{"CFullName":"組織行為","CName":"組行","EFullName":"Organizational Behavior","EName":"OB","pre_course":["Mng"],"status":0,"color":"yellow","CDescription":"本課程將介紹與組織中的個人、團體之行為相關的理論，並輔以研究案例之詮釋。與個人行為相關之主題包含：個體差異、感知能力、態度、學習能力、個人特質、決策與激勵動機；與團體組織行為相關之主題包含：領導能力、團體決策、團體動力、團隊合作、組織文化、溝通、衝突管理與壓力管理。","EDescription":"This course introduces the knowledge of the behavior of individuals and groups in organizations with major theoretical positions and research findings are discussed. The specific topics regarding individual behaviors include individual differences, perception, attitudes, learning, personality, decision making, motivation and the topics regarding group and organization behaviors include leadership, group decision making, group dynamics, teamwork, organizational culture, communication, conflict management and stress management."},
	"MtM":{"CFullName":"行銷管理","CName":"行管","EFullName":"Marketing Management","EName":"MtM","pre_course":["Mng"],"status":0,"color":"yellow","CDescription":"本課程將同時以學術研究面向與實用管理面向討論行銷，涵蓋目標如下：(一)提供學生有關行銷觀念及行銷步驟的基本知識；(二)經由評估行銷環境及研究消費者行為來了解市場及消費者；(三)重點討論行銷策略及行銷組合的規劃與實施。此外，亦會介紹各種行銷延伸議題，如：電子商務、全球行銷、行銷倫理等。","EDescription":"This course is designed to provide an informed appreciation of marketing as an academic subject as well as an important management practice. Linked to this attempt are the three following aims: i) To provide students with fundamental knowledge of marketing concepts and the marketing process, and its role in relation to company strategy. ii) To understand the marketplace and consumers by examine the characteristics of volatile marketing environment and exploring customer buying behavior. iii) To focus on the planning and implementation of a marketing strategy and marketing mix. In additions, extending marketing issues including e-commerce, global marketing and marketing ethics will also be covered."},
	"OM":{"CFullName":"作業管理","CName":"作管","EFullName":"Operations Management","EName":"OM","pre_course":["OR"],"status":0,"color":"yellow","CDescription":"本課程的重點在介紹一切與作業(operations)有關的經營管理問題，並探討各種機能應如何充分配合與發展，以便能有效且充分應用人力、材物料及設備，進而生產或提供適質、適量且合乎消費者需求的產品/服務之各種活動。課程內容包括︰預測、作業策略、製程管理與設備配置、品質管理、產品/服務之設計、產能規劃、供應鏈管理、存貨管理及排程、專案管理等單元。並將實際學習服務活動相關的調查分析。 ","EDescription":"This class is designed to introduce students to the important topic of operations management with apprehension such as product and service design, capacity planning, management of quality and quality control, inventory planning, etc. The students in this class will learn not only the algorithms and techniques used to solve these related problems, but also the real-world applications that can comprehend these methods. The students are encouraged to utilize computers in every respect of this class."},
	"IMFP":{"CFullName":"資管新生專題","CName":"新生專題","EFullName":"IM Freshman Project","EName":"IMFP","pre_course":[],"status":0,"color":"blue","CDescription":"本課程透過學者與專家之經驗分享、企業參訪、分組討論、心得撰寫等方式，讓學生了解資管專業內涵、大學課程內容、未來職涯發展等。","EDescription":""},
	"PoIM1":{"CFullName":"資管專題一","CName":"專題一","EFullName":"Project on Information Management (I)","EName":"PoIM1","pre_course":[],"status":0,"color":"blue","CDescription":"本課程將依據開課教授之安排，進行不同領域之理論研究及實作應用。","EDescription":"This course will provide diverse theoreticl readings and practical implementations vary from the instructors expertise."},
	"PoIM2":{"CFullName":"資管專題二","CName":"專題二","EFullName":"Project on Information Management (Ⅱ)","EName":"PoIM2","pre_course":["PoIM1"],"status":0,"color":"blue","CDescription":"本課程將依據開課教授之安排，進行不同領域之理論研究及實作應用。","EDescription":"This course will provide diverse theoreticl readings and practical implementations vary from the instructors expertise."},
	"MPD1":{"CFullName":"資管專題討論一","CName":"專討一","EFullName":"Mis Panel Discussion (Ⅰ)","EName":"MPD1","pre_course":[],"status":0,"color":"blue","CDescription":"本課程將邀請資訊相關領域的專家學者、業界菁英來分享其知識與經驗，內容涵蓋資訊管理、資訊科技、資訊創新、企業管理等，使學生了解目前資訊管理領域的最新發展與實際應用。","EDescription":"This course will invite scholars and researchers in the IS and related fields as well as executives and managers from the industry to share their knowledge and experiences on issues related to information management, information technologies, innovation, entrepreneurship. This course aims to acquaint students with trends and/or practical developments in the field of MIS."},
	"MPD2":{"CFullName":"資管專題討論二","CName":"專討二","EFullName":"Mis Panel Discussion (Ⅱ)","EName":"MPD2","pre_course":[],"status":0,"color":"blue","CDescription":"本課程將邀請資訊相關領域的專家學者、業界菁英來分享其知識與經驗，內容涵蓋資訊管理、資訊科技、資訊創新、企業管理等，使學生了解目前資訊管理領域的最新發展與實際應用。","EDescription":"This course will invite scholars and researchers in the IS and related fields as well as executives and managers from the industry to share their knowledge and experiences on issues related to information management, information technologies, innovation, entrepreneurship. This course aims to acquaint students with trends and/or practical developments in the field of MIS."},
	"SILP":{"CFullName":"資訊法與資訊政策專題","CName":"資訊法","EFullName":"Seminar on Information Law and Policy","EName":"SILP","pre_course":[],"status":0,"color":"blue","CDescription":"本課程的主要目的，乃是希望提供一個衡平的分析架構，探討、資通科技法律與社會間的互動關係，深入瞭解科技發展可能衍生的社會及倫理爭議、以及法律如何在科技與社會演進過程中扮演適當的角色。","EDescription":""}
}
`
