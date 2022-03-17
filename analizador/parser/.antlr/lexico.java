// Generated from c:\Users\Marco\Proyectos\Visual Code\goland\db_rust\analizador\parser\lexico.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class lexico extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		S_PUNTO=1, S_COMA=2, S_PTCOMA=3, S_ASIGNAR=4, S_AMP=5, S_MATCH_OR=6, S_MATCH_RET=7, 
		S_MATCH_DEFAULT=8, S_SUMA=9, S_RESTA=10, S_POR=11, S_DIVISION=12, S_MODULO=13, 
		S_MAYOR=14, S_MENOR=15, S_MAYORI=16, S_MENORI=17, S_IGUAL=18, S_DIFERENTE=19, 
		S_OR=20, S_AND=21, S_MOD_RET=22, S_APAR=23, S_CPAR=24, S_DOSPT=25, S_ALIAS=26, 
		S_ACOR=27, S_CCOR=28, S_ALLAV=29, S_CLLAV=30, S_RANGO=31, R_TRUE=32, R_FALSE=33, 
		R_LET=34, R_MUT=35, R_INT=36, R_FLOAT=37, R_BOOL=38, R_CHAR=39, R_STRING=40, 
		R_STR=41, R_FN=42, R_MAIN=43, R_AS=44, R_TO_OWNED=45, R_TO_STRING=46, 
		R_POW=47, R_POWF=48, R_PRINTLN=49, R_ABS=50, R_SQRT=51, R_CLONE=52, R_NEW=53, 
		R_LEN=54, R_PUSH=55, R_REMOVE=56, R_CONTAINS=57, R_INSERT=58, R_CAPACITY=59, 
		R_WITH_CAPACITY=60, R_IF=61, R_ELSE=62, R_MATCH=63, R_LOOP=64, R_BREAK=65, 
		R_WHILE=66, R_FOR=67, R_IN=68, R_CONTINUE=69, R_RETURN=70, R_STRUCT=71, 
		R_VECTOR=72, R_VEC=73, R_MOD=74, R_PUB=75, NUMERO=76, DECIMAL=77, CARACTER=78, 
		CADENA=79, ID=80, COMENTARIO=81, BLANCOS=82;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"S_PUNTO", "S_COMA", "S_PTCOMA", "S_ASIGNAR", "S_AMP", "S_MATCH_OR", 
			"S_MATCH_RET", "S_MATCH_DEFAULT", "S_SUMA", "S_RESTA", "S_POR", "S_DIVISION", 
			"S_MODULO", "S_MAYOR", "S_MENOR", "S_MAYORI", "S_MENORI", "S_IGUAL", 
			"S_DIFERENTE", "S_OR", "S_AND", "S_MOD_RET", "S_APAR", "S_CPAR", "S_DOSPT", 
			"S_ALIAS", "S_ACOR", "S_CCOR", "S_ALLAV", "S_CLLAV", "S_RANGO", "R_TRUE", 
			"R_FALSE", "R_LET", "R_MUT", "R_INT", "R_FLOAT", "R_BOOL", "R_CHAR", 
			"R_STRING", "R_STR", "R_FN", "R_MAIN", "R_AS", "R_TO_OWNED", "R_TO_STRING", 
			"R_POW", "R_POWF", "R_PRINTLN", "R_ABS", "R_SQRT", "R_CLONE", "R_NEW", 
			"R_LEN", "R_PUSH", "R_REMOVE", "R_CONTAINS", "R_INSERT", "R_CAPACITY", 
			"R_WITH_CAPACITY", "R_IF", "R_ELSE", "R_MATCH", "R_LOOP", "R_BREAK", 
			"R_WHILE", "R_FOR", "R_IN", "R_CONTINUE", "R_RETURN", "R_STRUCT", "R_VECTOR", 
			"R_VEC", "R_MOD", "R_PUB", "NUMERO", "DECIMAL", "CARACTER", "CADENA", 
			"ID", "COMENTARIO", "BLANCOS", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "','", "';'", "'='", "'&'", "'|'", "'=>'", "'_'", "'+'", 
			"'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", 
			"'||'", "'&&'", "'->'", "'('", "')'", "':'", "'::'", "'['", "']'", "'{'", 
			"'}'", "'..'", "'true'", "'false'", "'let'", "'mut'", "'i64'", "'f64'", 
			"'bool'", "'char'", "'string'", "'&str'", "'fn'", "'main'", "'as'", "'to_owned'", 
			"'to_string'", "'pow'", "'powF'", "'println!'", "'abs'", "'sqrt'", "'clone'", 
			"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'", 
			"'with_capacity'", "'if'", "'else'", "'match'", "'loop'", "'break'", 
			"'while'", "'for'", "'in'", "'continue'", "'return'", "'struct'", "'vector'", 
			"'vec'", "'mod'", "'pub'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "S_PUNTO", "S_COMA", "S_PTCOMA", "S_ASIGNAR", "S_AMP", "S_MATCH_OR", 
			"S_MATCH_RET", "S_MATCH_DEFAULT", "S_SUMA", "S_RESTA", "S_POR", "S_DIVISION", 
			"S_MODULO", "S_MAYOR", "S_MENOR", "S_MAYORI", "S_MENORI", "S_IGUAL", 
			"S_DIFERENTE", "S_OR", "S_AND", "S_MOD_RET", "S_APAR", "S_CPAR", "S_DOSPT", 
			"S_ALIAS", "S_ACOR", "S_CCOR", "S_ALLAV", "S_CLLAV", "S_RANGO", "R_TRUE", 
			"R_FALSE", "R_LET", "R_MUT", "R_INT", "R_FLOAT", "R_BOOL", "R_CHAR", 
			"R_STRING", "R_STR", "R_FN", "R_MAIN", "R_AS", "R_TO_OWNED", "R_TO_STRING", 
			"R_POW", "R_POWF", "R_PRINTLN", "R_ABS", "R_SQRT", "R_CLONE", "R_NEW", 
			"R_LEN", "R_PUSH", "R_REMOVE", "R_CONTAINS", "R_INSERT", "R_CAPACITY", 
			"R_WITH_CAPACITY", "R_IF", "R_ELSE", "R_MATCH", "R_LOOP", "R_BREAK", 
			"R_WHILE", "R_FOR", "R_IN", "R_CONTINUE", "R_RETURN", "R_STRUCT", "R_VECTOR", 
			"R_VEC", "R_MOD", "R_PUB", "NUMERO", "DECIMAL", "CARACTER", "CADENA", 
			"ID", "COMENTARIO", "BLANCOS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public lexico(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "lexico.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2T\u0227\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3"+
		"\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3"+
		"\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3"+
		"\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3"+
		"\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3 \3!\3!\3!\3!\3!\3"+
		"\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3"+
		"&\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*"+
		"\3*\3+\3+\3+\3,\3,\3,\3,\3,\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3.\3.\3/\3/"+
		"\3/\3/\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61"+
		"\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63\3\64"+
		"\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66"+
		"\3\67\3\67\3\67\3\67\38\38\38\38\38\39\39\39\39\39\39\39\3:\3:\3:\3:\3"+
		":\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3<\3<\3<\3<\3=\3=\3"+
		"=\3=\3=\3=\3=\3=\3=\3=\3=\3=\3=\3=\3>\3>\3>\3?\3?\3?\3?\3?\3@\3@\3@\3"+
		"@\3@\3@\3A\3A\3A\3A\3A\3B\3B\3B\3B\3B\3B\3C\3C\3C\3C\3C\3C\3D\3D\3D\3"+
		"D\3E\3E\3E\3F\3F\3F\3F\3F\3F\3F\3F\3F\3G\3G\3G\3G\3G\3G\3G\3H\3H\3H\3"+
		"H\3H\3H\3H\3I\3I\3I\3I\3I\3I\3I\3J\3J\3J\3J\3K\3K\3K\3K\3L\3L\3L\3L\3"+
		"M\6M\u01ee\nM\rM\16M\u01ef\3N\6N\u01f3\nN\rN\16N\u01f4\3N\3N\6N\u01f9"+
		"\nN\rN\16N\u01fa\3O\3O\3O\3O\3P\3P\7P\u0203\nP\fP\16P\u0206\13P\3P\3P"+
		"\3Q\3Q\7Q\u020c\nQ\fQ\16Q\u020f\13Q\3R\3R\3R\3R\7R\u0215\nR\fR\16R\u0218"+
		"\13R\3R\3R\3R\3R\3S\6S\u021f\nS\rS\16S\u0220\3S\3S\3T\3T\3T\2\2U\3\3\5"+
		"\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21"+
		"!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!"+
		"A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s"+
		";u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008f"+
		"I\u0091J\u0093K\u0095L\u0097M\u0099N\u009bO\u009dP\u009fQ\u00a1R\u00a3"+
		"S\u00a5T\u00a7\2\3\2\f\3\2\62;\3\2\60\60\3\2))\3\2$$\5\2C\\aac|\6\2\62"+
		";C\\aac|\4\2\60\60``\3\2\f\f\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2"+
		"\u022c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2"+
		"\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3"+
		"\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2"+
		"\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2"+
		"/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2"+
		"\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2"+
		"G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3"+
		"\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2"+
		"\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2"+
		"m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3"+
		"\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2"+
		"\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2"+
		"\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095"+
		"\3\2\2\2\2\u0097\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2\2\2\u009d\3\2\2"+
		"\2\2\u009f\3\2\2\2\2\u00a1\3\2\2\2\2\u00a3\3\2\2\2\2\u00a5\3\2\2\2\3\u00a9"+
		"\3\2\2\2\5\u00ab\3\2\2\2\7\u00ad\3\2\2\2\t\u00af\3\2\2\2\13\u00b1\3\2"+
		"\2\2\r\u00b3\3\2\2\2\17\u00b5\3\2\2\2\21\u00b8\3\2\2\2\23\u00ba\3\2\2"+
		"\2\25\u00bc\3\2\2\2\27\u00be\3\2\2\2\31\u00c0\3\2\2\2\33\u00c2\3\2\2\2"+
		"\35\u00c4\3\2\2\2\37\u00c6\3\2\2\2!\u00c8\3\2\2\2#\u00cb\3\2\2\2%\u00ce"+
		"\3\2\2\2\'\u00d1\3\2\2\2)\u00d4\3\2\2\2+\u00d7\3\2\2\2-\u00da\3\2\2\2"+
		"/\u00dd\3\2\2\2\61\u00df\3\2\2\2\63\u00e1\3\2\2\2\65\u00e3\3\2\2\2\67"+
		"\u00e6\3\2\2\29\u00e8\3\2\2\2;\u00ea\3\2\2\2=\u00ec\3\2\2\2?\u00ee\3\2"+
		"\2\2A\u00f1\3\2\2\2C\u00f6\3\2\2\2E\u00fc\3\2\2\2G\u0100\3\2\2\2I\u0104"+
		"\3\2\2\2K\u0108\3\2\2\2M\u010c\3\2\2\2O\u0111\3\2\2\2Q\u0116\3\2\2\2S"+
		"\u011d\3\2\2\2U\u0122\3\2\2\2W\u0125\3\2\2\2Y\u012a\3\2\2\2[\u012d\3\2"+
		"\2\2]\u0136\3\2\2\2_\u0140\3\2\2\2a\u0144\3\2\2\2c\u0149\3\2\2\2e\u0152"+
		"\3\2\2\2g\u0156\3\2\2\2i\u015b\3\2\2\2k\u0161\3\2\2\2m\u0165\3\2\2\2o"+
		"\u0169\3\2\2\2q\u016e\3\2\2\2s\u0175\3\2\2\2u\u017e\3\2\2\2w\u0185\3\2"+
		"\2\2y\u018e\3\2\2\2{\u019c\3\2\2\2}\u019f\3\2\2\2\177\u01a4\3\2\2\2\u0081"+
		"\u01aa\3\2\2\2\u0083\u01af\3\2\2\2\u0085\u01b5\3\2\2\2\u0087\u01bb\3\2"+
		"\2\2\u0089\u01bf\3\2\2\2\u008b\u01c2\3\2\2\2\u008d\u01cb\3\2\2\2\u008f"+
		"\u01d2\3\2\2\2\u0091\u01d9\3\2\2\2\u0093\u01e0\3\2\2\2\u0095\u01e4\3\2"+
		"\2\2\u0097\u01e8\3\2\2\2\u0099\u01ed\3\2\2\2\u009b\u01f2\3\2\2\2\u009d"+
		"\u01fc\3\2\2\2\u009f\u0200\3\2\2\2\u00a1\u0209\3\2\2\2\u00a3\u0210\3\2"+
		"\2\2\u00a5\u021e\3\2\2\2\u00a7\u0224\3\2\2\2\u00a9\u00aa\7\60\2\2\u00aa"+
		"\4\3\2\2\2\u00ab\u00ac\7.\2\2\u00ac\6\3\2\2\2\u00ad\u00ae\7=\2\2\u00ae"+
		"\b\3\2\2\2\u00af\u00b0\7?\2\2\u00b0\n\3\2\2\2\u00b1\u00b2\7(\2\2\u00b2"+
		"\f\3\2\2\2\u00b3\u00b4\7~\2\2\u00b4\16\3\2\2\2\u00b5\u00b6\7?\2\2\u00b6"+
		"\u00b7\7@\2\2\u00b7\20\3\2\2\2\u00b8\u00b9\7a\2\2\u00b9\22\3\2\2\2\u00ba"+
		"\u00bb\7-\2\2\u00bb\24\3\2\2\2\u00bc\u00bd\7/\2\2\u00bd\26\3\2\2\2\u00be"+
		"\u00bf\7,\2\2\u00bf\30\3\2\2\2\u00c0\u00c1\7\61\2\2\u00c1\32\3\2\2\2\u00c2"+
		"\u00c3\7\'\2\2\u00c3\34\3\2\2\2\u00c4\u00c5\7@\2\2\u00c5\36\3\2\2\2\u00c6"+
		"\u00c7\7>\2\2\u00c7 \3\2\2\2\u00c8\u00c9\7@\2\2\u00c9\u00ca\7?\2\2\u00ca"+
		"\"\3\2\2\2\u00cb\u00cc\7>\2\2\u00cc\u00cd\7?\2\2\u00cd$\3\2\2\2\u00ce"+
		"\u00cf\7?\2\2\u00cf\u00d0\7?\2\2\u00d0&\3\2\2\2\u00d1\u00d2\7#\2\2\u00d2"+
		"\u00d3\7?\2\2\u00d3(\3\2\2\2\u00d4\u00d5\7~\2\2\u00d5\u00d6\7~\2\2\u00d6"+
		"*\3\2\2\2\u00d7\u00d8\7(\2\2\u00d8\u00d9\7(\2\2\u00d9,\3\2\2\2\u00da\u00db"+
		"\7/\2\2\u00db\u00dc\7@\2\2\u00dc.\3\2\2\2\u00dd\u00de\7*\2\2\u00de\60"+
		"\3\2\2\2\u00df\u00e0\7+\2\2\u00e0\62\3\2\2\2\u00e1\u00e2\7<\2\2\u00e2"+
		"\64\3\2\2\2\u00e3\u00e4\7<\2\2\u00e4\u00e5\7<\2\2\u00e5\66\3\2\2\2\u00e6"+
		"\u00e7\7]\2\2\u00e78\3\2\2\2\u00e8\u00e9\7_\2\2\u00e9:\3\2\2\2\u00ea\u00eb"+
		"\7}\2\2\u00eb<\3\2\2\2\u00ec\u00ed\7\177\2\2\u00ed>\3\2\2\2\u00ee\u00ef"+
		"\7\60\2\2\u00ef\u00f0\7\60\2\2\u00f0@\3\2\2\2\u00f1\u00f2\7v\2\2\u00f2"+
		"\u00f3\7t\2\2\u00f3\u00f4\7w\2\2\u00f4\u00f5\7g\2\2\u00f5B\3\2\2\2\u00f6"+
		"\u00f7\7h\2\2\u00f7\u00f8\7c\2\2\u00f8\u00f9\7n\2\2\u00f9\u00fa\7u\2\2"+
		"\u00fa\u00fb\7g\2\2\u00fbD\3\2\2\2\u00fc\u00fd\7n\2\2\u00fd\u00fe\7g\2"+
		"\2\u00fe\u00ff\7v\2\2\u00ffF\3\2\2\2\u0100\u0101\7o\2\2\u0101\u0102\7"+
		"w\2\2\u0102\u0103\7v\2\2\u0103H\3\2\2\2\u0104\u0105\7k\2\2\u0105\u0106"+
		"\78\2\2\u0106\u0107\7\66\2\2\u0107J\3\2\2\2\u0108\u0109\7h\2\2\u0109\u010a"+
		"\78\2\2\u010a\u010b\7\66\2\2\u010bL\3\2\2\2\u010c\u010d\7d\2\2\u010d\u010e"+
		"\7q\2\2\u010e\u010f\7q\2\2\u010f\u0110\7n\2\2\u0110N\3\2\2\2\u0111\u0112"+
		"\7e\2\2\u0112\u0113\7j\2\2\u0113\u0114\7c\2\2\u0114\u0115\7t\2\2\u0115"+
		"P\3\2\2\2\u0116\u0117\7u\2\2\u0117\u0118\7v\2\2\u0118\u0119\7t\2\2\u0119"+
		"\u011a\7k\2\2\u011a\u011b\7p\2\2\u011b\u011c\7i\2\2\u011cR\3\2\2\2\u011d"+
		"\u011e\7(\2\2\u011e\u011f\7u\2\2\u011f\u0120\7v\2\2\u0120\u0121\7t\2\2"+
		"\u0121T\3\2\2\2\u0122\u0123\7h\2\2\u0123\u0124\7p\2\2\u0124V\3\2\2\2\u0125"+
		"\u0126\7o\2\2\u0126\u0127\7c\2\2\u0127\u0128\7k\2\2\u0128\u0129\7p\2\2"+
		"\u0129X\3\2\2\2\u012a\u012b\7c\2\2\u012b\u012c\7u\2\2\u012cZ\3\2\2\2\u012d"+
		"\u012e\7v\2\2\u012e\u012f\7q\2\2\u012f\u0130\7a\2\2\u0130\u0131\7q\2\2"+
		"\u0131\u0132\7y\2\2\u0132\u0133\7p\2\2\u0133\u0134\7g\2\2\u0134\u0135"+
		"\7f\2\2\u0135\\\3\2\2\2\u0136\u0137\7v\2\2\u0137\u0138\7q\2\2\u0138\u0139"+
		"\7a\2\2\u0139\u013a\7u\2\2\u013a\u013b\7v\2\2\u013b\u013c\7t\2\2\u013c"+
		"\u013d\7k\2\2\u013d\u013e\7p\2\2\u013e\u013f\7i\2\2\u013f^\3\2\2\2\u0140"+
		"\u0141\7r\2\2\u0141\u0142\7q\2\2\u0142\u0143\7y\2\2\u0143`\3\2\2\2\u0144"+
		"\u0145\7r\2\2\u0145\u0146\7q\2\2\u0146\u0147\7y\2\2\u0147\u0148\7H\2\2"+
		"\u0148b\3\2\2\2\u0149\u014a\7r\2\2\u014a\u014b\7t\2\2\u014b\u014c\7k\2"+
		"\2\u014c\u014d\7p\2\2\u014d\u014e\7v\2\2\u014e\u014f\7n\2\2\u014f\u0150"+
		"\7p\2\2\u0150\u0151\7#\2\2\u0151d\3\2\2\2\u0152\u0153\7c\2\2\u0153\u0154"+
		"\7d\2\2\u0154\u0155\7u\2\2\u0155f\3\2\2\2\u0156\u0157\7u\2\2\u0157\u0158"+
		"\7s\2\2\u0158\u0159\7t\2\2\u0159\u015a\7v\2\2\u015ah\3\2\2\2\u015b\u015c"+
		"\7e\2\2\u015c\u015d\7n\2\2\u015d\u015e\7q\2\2\u015e\u015f\7p\2\2\u015f"+
		"\u0160\7g\2\2\u0160j\3\2\2\2\u0161\u0162\7p\2\2\u0162\u0163\7g\2\2\u0163"+
		"\u0164\7y\2\2\u0164l\3\2\2\2\u0165\u0166\7n\2\2\u0166\u0167\7g\2\2\u0167"+
		"\u0168\7p\2\2\u0168n\3\2\2\2\u0169\u016a\7r\2\2\u016a\u016b\7w\2\2\u016b"+
		"\u016c\7u\2\2\u016c\u016d\7j\2\2\u016dp\3\2\2\2\u016e\u016f\7t\2\2\u016f"+
		"\u0170\7g\2\2\u0170\u0171\7o\2\2\u0171\u0172\7q\2\2\u0172\u0173\7x\2\2"+
		"\u0173\u0174\7g\2\2\u0174r\3\2\2\2\u0175\u0176\7e\2\2\u0176\u0177\7q\2"+
		"\2\u0177\u0178\7p\2\2\u0178\u0179\7v\2\2\u0179\u017a\7c\2\2\u017a\u017b"+
		"\7k\2\2\u017b\u017c\7p\2\2\u017c\u017d\7u\2\2\u017dt\3\2\2\2\u017e\u017f"+
		"\7k\2\2\u017f\u0180\7p\2\2\u0180\u0181\7u\2\2\u0181\u0182\7g\2\2\u0182"+
		"\u0183\7t\2\2\u0183\u0184\7v\2\2\u0184v\3\2\2\2\u0185\u0186\7e\2\2\u0186"+
		"\u0187\7c\2\2\u0187\u0188\7r\2\2\u0188\u0189\7c\2\2\u0189\u018a\7e\2\2"+
		"\u018a\u018b\7k\2\2\u018b\u018c\7v\2\2\u018c\u018d\7{\2\2\u018dx\3\2\2"+
		"\2\u018e\u018f\7y\2\2\u018f\u0190\7k\2\2\u0190\u0191\7v\2\2\u0191\u0192"+
		"\7j\2\2\u0192\u0193\7a\2\2\u0193\u0194\7e\2\2\u0194\u0195\7c\2\2\u0195"+
		"\u0196\7r\2\2\u0196\u0197\7c\2\2\u0197\u0198\7e\2\2\u0198\u0199\7k\2\2"+
		"\u0199\u019a\7v\2\2\u019a\u019b\7{\2\2\u019bz\3\2\2\2\u019c\u019d\7k\2"+
		"\2\u019d\u019e\7h\2\2\u019e|\3\2\2\2\u019f\u01a0\7g\2\2\u01a0\u01a1\7"+
		"n\2\2\u01a1\u01a2\7u\2\2\u01a2\u01a3\7g\2\2\u01a3~\3\2\2\2\u01a4\u01a5"+
		"\7o\2\2\u01a5\u01a6\7c\2\2\u01a6\u01a7\7v\2\2\u01a7\u01a8\7e\2\2\u01a8"+
		"\u01a9\7j\2\2\u01a9\u0080\3\2\2\2\u01aa\u01ab\7n\2\2\u01ab\u01ac\7q\2"+
		"\2\u01ac\u01ad\7q\2\2\u01ad\u01ae\7r\2\2\u01ae\u0082\3\2\2\2\u01af\u01b0"+
		"\7d\2\2\u01b0\u01b1\7t\2\2\u01b1\u01b2\7g\2\2\u01b2\u01b3\7c\2\2\u01b3"+
		"\u01b4\7m\2\2\u01b4\u0084\3\2\2\2\u01b5\u01b6\7y\2\2\u01b6\u01b7\7j\2"+
		"\2\u01b7\u01b8\7k\2\2\u01b8\u01b9\7n\2\2\u01b9\u01ba\7g\2\2\u01ba\u0086"+
		"\3\2\2\2\u01bb\u01bc\7h\2\2\u01bc\u01bd\7q\2\2\u01bd\u01be\7t\2\2\u01be"+
		"\u0088\3\2\2\2\u01bf\u01c0\7k\2\2\u01c0\u01c1\7p\2\2\u01c1\u008a\3\2\2"+
		"\2\u01c2\u01c3\7e\2\2\u01c3\u01c4\7q\2\2\u01c4\u01c5\7p\2\2\u01c5\u01c6"+
		"\7v\2\2\u01c6\u01c7\7k\2\2\u01c7\u01c8\7p\2\2\u01c8\u01c9\7w\2\2\u01c9"+
		"\u01ca\7g\2\2\u01ca\u008c\3\2\2\2\u01cb\u01cc\7t\2\2\u01cc\u01cd\7g\2"+
		"\2\u01cd\u01ce\7v\2\2\u01ce\u01cf\7w\2\2\u01cf\u01d0\7t\2\2\u01d0\u01d1"+
		"\7p\2\2\u01d1\u008e\3\2\2\2\u01d2\u01d3\7u\2\2\u01d3\u01d4\7v\2\2\u01d4"+
		"\u01d5\7t\2\2\u01d5\u01d6\7w\2\2\u01d6\u01d7\7e\2\2\u01d7\u01d8\7v\2\2"+
		"\u01d8\u0090\3\2\2\2\u01d9\u01da\7x\2\2\u01da\u01db\7g\2\2\u01db\u01dc"+
		"\7e\2\2\u01dc\u01dd\7v\2\2\u01dd\u01de\7q\2\2\u01de\u01df\7t\2\2\u01df"+
		"\u0092\3\2\2\2\u01e0\u01e1\7x\2\2\u01e1\u01e2\7g\2\2\u01e2\u01e3\7e\2"+
		"\2\u01e3\u0094\3\2\2\2\u01e4\u01e5\7o\2\2\u01e5\u01e6\7q\2\2\u01e6\u01e7"+
		"\7f\2\2\u01e7\u0096\3\2\2\2\u01e8\u01e9\7r\2\2\u01e9\u01ea\7w\2\2\u01ea"+
		"\u01eb\7d\2\2\u01eb\u0098\3\2\2\2\u01ec\u01ee\t\2\2\2\u01ed\u01ec\3\2"+
		"\2\2\u01ee\u01ef\3\2\2\2\u01ef\u01ed\3\2\2\2\u01ef\u01f0\3\2\2\2\u01f0"+
		"\u009a\3\2\2\2\u01f1\u01f3\t\2\2\2\u01f2\u01f1\3\2\2\2\u01f3\u01f4\3\2"+
		"\2\2\u01f4\u01f2\3\2\2\2\u01f4\u01f5\3\2\2\2\u01f5\u01f6\3\2\2\2\u01f6"+
		"\u01f8\t\3\2\2\u01f7\u01f9\t\2\2\2\u01f8\u01f7\3\2\2\2\u01f9\u01fa\3\2"+
		"\2\2\u01fa\u01f8\3\2\2\2\u01fa\u01fb\3\2\2\2\u01fb\u009c\3\2\2\2\u01fc"+
		"\u01fd\t\4\2\2\u01fd\u01fe\13\2\2\2\u01fe\u01ff\t\4\2\2\u01ff\u009e\3"+
		"\2\2\2\u0200\u0204\7$\2\2\u0201\u0203\n\5\2\2\u0202\u0201\3\2\2\2\u0203"+
		"\u0206\3\2\2\2\u0204\u0202\3\2\2\2\u0204\u0205\3\2\2\2\u0205\u0207\3\2"+
		"\2\2\u0206\u0204\3\2\2\2\u0207\u0208\7$\2\2\u0208\u00a0\3\2\2\2\u0209"+
		"\u020d\t\6\2\2\u020a\u020c\t\7\2\2\u020b\u020a\3\2\2\2\u020c\u020f\3\2"+
		"\2\2\u020d\u020b\3\2\2\2\u020d\u020e\3\2\2\2\u020e\u00a2\3\2\2\2\u020f"+
		"\u020d\3\2\2\2\u0210\u0211\7\61\2\2\u0211\u0212\7\61\2\2\u0212\u0216\3"+
		"\2\2\2\u0213\u0215\t\b\2\2\u0214\u0213\3\2\2\2\u0215\u0218\3\2\2\2\u0216"+
		"\u0214\3\2\2\2\u0216\u0217\3\2\2\2\u0217\u0219\3\2\2\2\u0218\u0216\3\2"+
		"\2\2\u0219\u021a\t\t\2\2\u021a\u021b\3\2\2\2\u021b\u021c\bR\2\2\u021c"+
		"\u00a4\3\2\2\2\u021d\u021f\t\n\2\2\u021e\u021d\3\2\2\2\u021f\u0220\3\2"+
		"\2\2\u0220\u021e\3\2\2\2\u0220\u0221\3\2\2\2\u0221\u0222\3\2\2\2\u0222"+
		"\u0223\bS\2\2\u0223\u00a6\3\2\2\2\u0224\u0225\7^\2\2\u0225\u0226\t\13"+
		"\2\2\u0226\u00a8\3\2\2\2\n\2\u01ef\u01f4\u01fa\u0204\u020d\u0216\u0220"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}