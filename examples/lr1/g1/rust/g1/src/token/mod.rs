
//! Module token is generated by GoGLL. Do not edit

extern crate lazy_static;

use std::rc::Rc;
use std::fmt;
use lazy_static::lazy_static;
use std::collections::HashMap;

/// Token is returned by the lexer for every scanned lexical token
pub struct Token {
	pub typ: Type,
	pub lext: usize, 
	pub rext: usize,
	input: Rc<Vec<char>>,
}

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
pub enum Type {	
	Error, // "Error"
	EOF, // "EOF"
	Type0, // "+"
	Type1, // "a"
}

/**
New returns a new token.  
lext is the left extent and rext the right extent of the token in the input.  
input is the input slice scanned by the lexer.
*/
pub fn new<'a>(t: Type, lext: usize, rext: usize, input: &Rc<Vec<char>>) -> Rc<Token> {
	Rc::new(Token{
		typ:   t,
		lext:  lext,
		rext:  rext,
		input: input.clone(),
	})
}

impl Token {
	/// get_line_column returns the (line, column) of the left extent of the token
	pub fn get_line_column(&self) -> (usize, usize) {
		let mut line = 1;
		let mut col = 1;
		let mut j = 0;
		while j < self.lext {
			match self.input[j] {
			'\n' => {
				line += 1;
				col = 1
			},
			'\t' => col += 4,
			_ => col += 1
			}
			j += 1
		}
		(line, col)
	}
	
	/// literal returns the literal runes of t scanned by the lexer
	pub fn literal(&self) -> Vec<char> {
		self.input[self.lext..self.rext].to_vec()
	}
	
    /// literal_string returs the literal string of t scanned by the lexer
    #[allow(dead_code)]
	pub fn literal_string(&self) -> String {
		self.literal().iter().collect::<String>()
    }

} // impl Token

impl <'a>fmt::Display for Token {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		let (ln, col) = self.get_line_column();
		write!(f, "({}, ({},{}) {})", 
			self.typ, ln, col, self.literal().iter().collect::<String>())
	}

}

impl <'a>Type {
	/// id returns the token type ID of token Type t
	#[allow(dead_code)]
	pub fn id(&self) -> &'a str {
		TYPE_TO_STRING[self]
	}
	
}

impl <'a>fmt::Display for Type {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		write!(f, "{}", TYPE_TO_STRING[self])
	}

}

lazy_static! {
    static ref TYPE_TO_STRING: HashMap<Type, &'static str> = {
        let mut m = HashMap::new(); 
		m.insert(Type::Error, "Error");
		m.insert(Type::EOF, "EOF");
		m.insert(Type::Type0, "+");
		m.insert(Type::Type1, "a");
        m
    };
}

lazy_static! {
	static ref STRING_TO_TYPE: HashMap<&'static str, Type> = {
		let mut m = HashMap::new(); 
		m.insert("Error", Type::Error); 
		m.insert("EOF", Type::EOF); 
		m.insert("Type0", Type::Type0); 
		m.insert("Type1", Type::Type1); 
		m
	};
}
