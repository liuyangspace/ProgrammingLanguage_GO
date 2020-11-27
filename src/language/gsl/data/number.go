package data


/*

byte
int
	int_lit     = decimal_lit | octal_lit | hex_lit .
	decimal_lit = ( "1" … "9" ) { decimal_digit } .
	octal_lit   = "0" { octal_digit } .
	hex_lit     = "0" ( "x" | "X" ) hex_digit { hex_digit } .
	e.g:
		42
		0600
		0xBadFace
		170141183460469231731687303715884105727
float
	float_lit = decimals "." [ decimals ] [ exponent ] | decimals exponent |             "."
	decimals [ exponent ] . decimals  = decimal_digit { decimal_digit } .
	exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals .
	e.g:
		0.
		72.40
		072.40  // == 72.40
		2.71828
		1.e+0
		6.67428e-11
		1E6
		.25
		.12345E+5
 complex
	imaginary_lit = (decimals | float_lit) "i" .
	e.g:
		0i
		011i  // == 11i
		0.i
		2.71828i
		1.e+0i
		6.67428e-11i
		1E6i
		.25i
		.12345E+5i

 */
