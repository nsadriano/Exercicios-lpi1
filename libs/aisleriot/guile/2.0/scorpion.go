GOOF----LE-8-2.0(      ] H 4      h(      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤							 ¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-extended-slot¤	g  down¤	g  
deal-cards¤				 ¤	g  deal-cards-face-up¤				 ¤	g  check-score¤	g  new-game¤	g  is-visible?¤	g  check-score-cards¤	 g  get-suit¤	!g  	get-value¤	"g  check-score-slot¤	#g  	get-cards¤	$g  
set-score!¤	%g  empty-slot?¤	&g  reverse¤	'g  button-pressed¤	(g  length¤	)g  correct-sequence¤	*g  king¤	+g  get-top-card¤	,g  
droppable?¤	-g  move-n-cards!¤	.g  make-visible-top-card¤	/g  button-released¤	0		 ¤	1g  button-clicked¤	2g  button-double-clicked¤	3g  get-hint¤	4g  game-continuable¤	5g  slots-filled?¤	6g  game-won¤	7g  _¤	8f  Deal the cards¤	9g  	dealable?¤	:g  check-slot-cards¤	;g  	hint-move¤	<g  check-a-slot¤	=g  
check-slot¤	>g  find-empty-slot¤	?g  here-kingy-kingy¤	@g  king-avail?¤	Ag  check-for-empty¤	Bg  get-options¤	Cg  apply-options¤	Dg  timeout¤	Eg  set-features¤	Fg  droppable-feature¤	Gg  
set-lambda¤C 5   hP!  -  ] 4	 >  "  G  R        hÀ  \  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4

>  "  G  4
>  "  G  4

>  "  G  4
>  "  G  4

>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  			 C T      g  filenamef  scorpion.scm
	
							#			3			C			U			e			h			o			x	 		{	 	 	 	 	!	 	!	 	!	 	"	 ¡	"	 ¨	"	 ±	#	 ´	#	 »	#	 Ä	$	 Ç	$	 Î	$	 ×	%	 Ú	%	 á	%	 ê	'	 ï	'	 ô	'	 ý	(		(		(		)		)		)	#	*	(	*	-	*	6	+	;	+	@	+	I	,	N	,	S	,	\	-	a	-	f	-	o	.	t	.	y	.		/		/		/		0		0		0	¨	2	¾	4	 =	¿
  g  nameg  new-game CR !     hÈ   Z  ](  $  	$   	C C C45$  $  "  
45$   6"   64545$  *4545$   6"ÿÿ³"ÿÿ¯ 	6    R      g  acc
	 Ä g  cards	 Ä g  unbroken		 Ä g  count		 Ä g  t		3	M  g  filenamef  scorpion.scm
	6
		7			9			9			9			:		$	<		)	<		+	<		/	7		2	>		3	>		3	>		A	?		F	?		I	?		J	?		Q	7		X	@		]	@	3	_	@		j	G		n	G	-	p	G		p	7		q	A		v	A		x	A		y	B		~	B	 	B	 	A	 	7	 	C	 	C	 	C	 	C	 	D	 	D	 	D	 	C	 	A	 ¢	E	 ¥	E	$ ª	E	9 ¬	E	 »	=	 ¾	=	$ Â	=	3 Ä	=	 3	 Ä	  g  nameg  check-score-cards CR"# h(   À   ](   C4 45
56     ¸       g  acc
		# g  slots		#  g  filenamef  scorpion.scm
	I
		J			L			L	/		L	:		L	/		L		!	L	M	#	L	 
		#	  g  nameg  check-score-slot C"R$"   h   f   ] 4	$56 ^       g  filenamef  scorpion.scm
	N
		O			O	 		
  g  nameg  check-score CR%&     h    ®   ]4 5$  C456     ¦       g  slot-id
		 g  	card-list		  g  filenamef  scorpion.scm
	Q
		R			R			S			S			S	 			  g  nameg  button-pressed C'R( !) h`   I  ]	4 5$  C4 5$  74 54 5&   4 54 5$   6CCC    A      g  	card-list
		\ g  t		\  g  filenamef  scorpion.scm
	U
		V				V			V			W			W			W		#	W		$	X		)	X		+	X		,	Y		1	Y		4	Y		8	W		9	Z		>	Z		@	Z		A	Z		B	[		G	[		J	[		K	Z		O	W		T	\		V	\	 		\  g  nameg  correct-sequence C)R%!&* +    h   «  ] $  C
$  C45$  4455"  $  C45$  C44554455&  44554455CC      £      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t		7   g  filenamef  scorpion.scm
	^
		_			_			`			_			a		"	a		#	b		&	b	#	-	b		/	b		2	b		7	a		C	c		M	c		P	d		S	d		[	d		\	e		_	e	$	f	e		h	e		l	c		m	f		p	f		x	f		y	g		|	g	( 	g	# 	g	 	g	 	f	 !	 	  g  nameg  
droppable? C,R,-%.    hP   ÷   ]4 5$  94 5$  (4 5$  "  4 5$  6 CCC    ï       g  
start-slot
		L g  	card-list		L g  end-slot			L g  t		'	>  g  filenamef  scorpion.scm
	i
		j			j			k		 	j		!	l		'	l		5	m		B	j		F	n	 		L	  g  nameg  button-released C/R%0  h0   ·   ] 
$  4
5$  C4
5$  6 CC      ¯       g  slot-id
		*  g  filenamef  scorpion.scm
	p
		q		
	q			r			q			s			s			s		"	q		&	t	 		*  g  nameg  button-clicked C1R    h   v   ]C    n       g  slot-id
		  g  filenamef  scorpion.scm
	v
 		  g  nameg  button-double-clicked C2R3 h   c   ] 6   [       g  filenamef  scorpion.scm
	y
		z	 		
  g  nameg  game-continuable C4R%5(#)    hP   /  ] (  C4 5$   6	44 55$  44 55$   6CC   '      g  slots
		M  g  filenamef  scorpion.scm
	|
		}		 		 		 			}		 		 		  		# 		( 	%	* 		, 		- 		1	}		2 		5 		: 	)	< 		> 		B 		G 		I 	 		M  g  nameg  slots-filled? C5R5      h   ]   ] 6U       g  filenamef  scorpion.scm
 
	 	 		
  g  nameg  game-won C6R%78      h       ] 4
5$  C
45 C             g  filenamef  scorpion.scm
 
	 		 		 		 		 		 	 		
  g  nameg  	dealable? C9R(: !;  hx   ¯  ]45
$  "  	45$  C"   64545&  #4545$  
 6"ÿÿº"ÿÿ¶§      g  slot1
		x g  slot2		x g  count			x g  card			x g  	card-list			x g  t			$  g  filenamef  scorpion.scm
 
	 		 		 			 		 		  		! 		( 		7 	*	< 	;	> 		> 		? 		F 		K 		M 		Q 		R 		Y 		^ 	!	` 		a 		b 		f 			p 		 		x	  g  nameg  check-slot-cards C:R%:+#<        hh     ]	$  C $  "  (45$  "  4 4 5455$   4 5456 6       g  slot1
		g g  slot2		g  g  filenamef  scorpion.scm
 
	 			 		 		 			 		& 			, 		4 	.	; 	C	C 		G 		O  	)	V  	>	^  			e ¡	 	g ¡	 		g	  g  nameg  check-a-slot C<R%<=    h8   ½   ] 	$  C4 5$  "  4 5$   6 6 µ       g  slot-id
		7  g  filenamef  scorpion.scm
 £
	 ¤			 ¤		 ¦		 ¦			 §		) ¤		0 ¨			5 ©		7 ©	 		7  g  nameg  
check-slot C=R(!*;>? 	   hp   w  ]45
$  "  "45$  "  	45$  C45$   456 6   o      g  slot-id
		m g  count		m g  	card-list			m g  t			= g  t		!	:  g  filenamef  scorpion.scm
 «
	 ¬		 ¬		 ¬			 ­		! ­		! ¬			/ ®		4 ®		6 ®		7 ®		A ¬		D °		I °		K °		N °			R ¬		Y ±	"	a ±			h ²	&	k ²	2	m ²	 		m	  g  nameg  here-kingy-kingy C?R%?#@  hH   Ù   ] 	$  C4 5$  "  4 4 55$   4 56 6   Ñ       g  slot-id
		E  g  filenamef  scorpion.scm
 ´
	 µ			 µ		 ·		 ·			 ¸		$ ¸	*	, ¸		0 µ		6 ¹	%	> ¹			C º		E º	 		E  g  nameg  king-avail? C@R%@    h    E  ]045  $  "  z4	5$  "  c4	5$  "  L4	5$  "  54	5$  "  4	5$  "  4	5 $  6C   =      g  t
	  g  t	  g  t		0  g  t		D  g  t		X  g  t		l   g  filenamef  scorpion.scm
 ¼
	 ½		 ½		 ¾		 ½		* ¿		0 ½		> À		D ½		R Á		X ½		f Â		l ½		z Ã	  ½	  Ä	 	 
  g  nameg  check-for-empty CAR=A9      h(      ]45  $   C45   $   C6         g  t
		' g  t
		'  g  filenamef  scorpion.scm
 Æ
	 Ç		 Ç		 È		 Ç		' É	 		'
  g  nameg  get-hint C3R    h   W   ] C    O       g  filenamef  scorpion.scm
 Ë
 		
  g  nameg  get-options CBR    h   o   ]C    g       g  options
		  g  filenamef  scorpion.scm
 Î
 		  g  nameg  apply-options CCR    h   S   ] C    K       g  filenamef  scorpion.scm
 Ñ
 		
  g  nameg  timeout CDR4EiFi>  "  G  Gii'i/i1i2i4i6i3iBiCiDi,i6      %      g  filenamef  scorpion.scm		
					
o	
­	6
«	I
	9	N

!	Q
ä	U
>	^
¢	i
¢	p
1	v
®	y
J	|
È 
 
Ò 
j 
x £
 «
¼ ´
¸ ¼
 Æ
 
 Ë
  Î
 þ Ñ
 ÿ Ô
!J Ö
 	!J
   C6 