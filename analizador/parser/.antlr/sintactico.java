// Generated from c:\Users\Marco\Proyectos\Visual Code\goland\db_rust\analizador\parser\sintactico.g4 by ANTLR 4.8

	import "db_rust/analizador/ast"
	import "db_rust/analizador/ast/expresion"
	import "db_rust/analizador/ast/funcion"
	import "db_rust/analizador/ast/imprimir"
	import "db_rust/analizador/ast/variables"
	import "db_rust/analizador/ast/interfaces"
	import "db_rust/analizador/entorno"
	import arrayList "github.com/colegno/arraylist"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class sintactico extends Parser {
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
	public static final int
		RULE_start = 0, RULE_funciones = 1, RULE_funcion = 2, RULE_funcMain = 3, 
		RULE_instrucciones = 4, RULE_instruccion = 5, RULE_imprimir = 6, RULE_lista_exp = 7, 
		RULE_declaracion = 8, RULE_tipo_dato = 9, RULE_exp = 10, RULE_logica = 11, 
		RULE_relacional = 12, RULE_aritmetica = 13, RULE_exp_valor = 14, RULE_primitivo = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funciones", "funcion", "funcMain", "instrucciones", "instruccion", 
			"imprimir", "lista_exp", "declaracion", "tipo_dato", "exp", "logica", 
			"relacional", "aritmetica", "exp_valor", "primitivo"
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

	@Override
	public String getGrammarFileName() { return "sintactico.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public sintactico(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ast.Ast root;
		public FuncionesContext funciones;
		public FuncionesContext funciones() {
			return getRuleContext(FuncionesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			((StartContext)_localctx).funciones = funciones();

					_localctx.root = ast.NewAst(((StartContext)_localctx).funciones.lista)
				
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public FuncionContext funcion;
		public List<FuncionContext> fun = new ArrayList<FuncionContext>();
		public List<FuncionContext> funcion() {
			return getRuleContexts(FuncionContext.class);
		}
		public FuncionContext funcion(int i) {
			return getRuleContext(FuncionContext.class,i);
		}
		public FuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciones; }
	}

	public final FuncionesContext funciones() throws RecognitionException {
		FuncionesContext _localctx = new FuncionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funciones);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==R_FN) {
				{
				{
				setState(35);
				((FuncionesContext)_localctx).funcion = funcion();
				((FuncionesContext)_localctx).fun.add(((FuncionesContext)_localctx).funcion);
				}
				}
				setState(40);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

					LISTA := localctx.(*FuncionesContext).GetFun()
					for _, i := range LISTA {
						_localctx.lista.Add(i.GetInstr())
					}
				
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public FuncMainContext funcMain;
		public FuncMainContext funcMain() {
			return getRuleContext(FuncMainContext.class,0);
		}
		public FuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcion; }
	}

	public final FuncionContext funcion() throws RecognitionException {
		FuncionContext _localctx = new FuncionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_funcion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			((FuncionContext)_localctx).funcMain = funcMain();
			 _localctx.instr = ((FuncionContext)_localctx).funcMain.instr 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncMainContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public InstruccionesContext instrucciones;
		public TerminalNode R_FN() { return getToken(sintactico.R_FN, 0); }
		public TerminalNode R_MAIN() { return getToken(sintactico.R_MAIN, 0); }
		public TerminalNode S_APAR() { return getToken(sintactico.S_APAR, 0); }
		public TerminalNode S_CPAR() { return getToken(sintactico.S_CPAR, 0); }
		public TerminalNode S_ALLAV() { return getToken(sintactico.S_ALLAV, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode S_CLLAV() { return getToken(sintactico.S_CLLAV, 0); }
		public FuncMainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcMain; }
	}

	public final FuncMainContext funcMain() throws RecognitionException {
		FuncMainContext _localctx = new FuncMainContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_funcMain);
		 params := arrayList.New() 
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(R_FN);
			setState(47);
			match(R_MAIN);
			setState(48);
			match(S_APAR);
			setState(49);
			match(S_CPAR);
			setState(50);
			match(S_ALLAV);
			setState(51);
			((FuncMainContext)_localctx).instrucciones = instrucciones();
			setState(52);
			match(S_CLLAV);

					_localctx.instr = funcion.NewFuncion(entorno.VOID, "main", params, ((FuncMainContext)_localctx).instrucciones.lista)
				
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionContext instruccion;
		public List<InstruccionContext> ins = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_instrucciones);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(55);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).ins.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(58); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==R_LET || _la==R_PRINTLN );

					LISTA := localctx.(*InstruccionesContext).GetIns()
					for _, i := range LISTA {
						_localctx.lista.Add(i.GetInstr())
					}
				
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ImprimirContext imprimir;
		public DeclaracionContext declaracion;
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public TerminalNode S_PTCOMA() { return getToken(sintactico.S_PTCOMA, 0); }
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_instruccion);
		try {
			setState(70);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_PRINTLN:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				((InstruccionContext)_localctx).imprimir = imprimir();
				setState(63);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).imprimir.instr 
				}
				break;
			case R_LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(67);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).declaracion.instr 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImprimirContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Lista_expContext lista_exp;
		public TerminalNode R_PRINTLN() { return getToken(sintactico.R_PRINTLN, 0); }
		public TerminalNode S_APAR() { return getToken(sintactico.S_APAR, 0); }
		public Lista_expContext lista_exp() {
			return getRuleContext(Lista_expContext.class,0);
		}
		public TerminalNode S_CPAR() { return getToken(sintactico.S_CPAR, 0); }
		public ImprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imprimir; }
	}

	public final ImprimirContext imprimir() throws RecognitionException {
		ImprimirContext _localctx = new ImprimirContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(R_PRINTLN);
			setState(73);
			match(S_APAR);
			setState(74);
			((ImprimirContext)_localctx).lista_exp = lista_exp(0);
			setState(75);
			match(S_CPAR);

					_localctx.instr = imprimir.NewImprimir(((ImprimirContext)_localctx).lista_exp.lista)
				
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Lista_expContext extends ParserRuleContext {
		public *arrayList.List lista;
		public Lista_expContext LISTA;
		public ExpContext exp;
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode S_COMA() { return getToken(sintactico.S_COMA, 0); }
		public Lista_expContext lista_exp() {
			return getRuleContext(Lista_expContext.class,0);
		}
		public Lista_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_exp; }
	}

	public final Lista_expContext lista_exp() throws RecognitionException {
		return lista_exp(0);
	}

	private Lista_expContext lista_exp(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Lista_expContext _localctx = new Lista_expContext(_ctx, _parentState);
		Lista_expContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_lista_exp, _p);
		 _localctx.lista = arrayList.New() 
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(79);
			((Lista_expContext)_localctx).exp = exp();
			 _localctx.lista.Add(((Lista_expContext)_localctx).exp.val) 
			}
			_ctx.stop = _input.LT(-1);
			setState(89);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Lista_expContext(_parentctx, _parentState);
					_localctx.LISTA = _prevctx;
					_localctx.LISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_lista_exp);
					setState(82);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(83);
					match(S_COMA);
					setState(84);
					((Lista_expContext)_localctx).exp = exp();
					 
					          		((Lista_expContext)_localctx).LISTA.lista.Add(((Lista_expContext)_localctx).exp.val) 
					          		_localctx.lista = ((Lista_expContext)_localctx).LISTA.lista 
					          	
					}
					} 
				}
				setState(91);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DeclaracionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token ID;
		public Tipo_datoContext tipo_dato;
		public ExpContext exp;
		public TerminalNode R_LET() { return getToken(sintactico.R_LET, 0); }
		public TerminalNode R_MUT() { return getToken(sintactico.R_MUT, 0); }
		public TerminalNode ID() { return getToken(sintactico.ID, 0); }
		public TerminalNode S_DOSPT() { return getToken(sintactico.S_DOSPT, 0); }
		public Tipo_datoContext tipo_dato() {
			return getRuleContext(Tipo_datoContext.class,0);
		}
		public TerminalNode S_ASIGNAR() { return getToken(sintactico.S_ASIGNAR, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declaracion);
		try {
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				match(R_LET);
				setState(93);
				match(R_MUT);
				setState(94);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(95);
				match(S_DOSPT);
				setState(96);
				((DeclaracionContext)_localctx).tipo_dato = tipo_dato();
				setState(97);
				match(S_ASIGNAR);
				setState(98);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(true, id, ((DeclaracionContext)_localctx).tipo_dato.tipo, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				match(R_LET);
				setState(102);
				match(R_MUT);
				setState(103);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(104);
				match(S_ASIGNAR);
				setState(105);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(true, id, entorno.NULL, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(108);
				match(R_LET);
				setState(109);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(110);
				match(S_DOSPT);
				setState(111);
				((DeclaracionContext)_localctx).tipo_dato = tipo_dato();
				setState(112);
				match(S_ASIGNAR);
				setState(113);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(false, id, ((DeclaracionContext)_localctx).tipo_dato.tipo, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(116);
				match(R_LET);
				setState(117);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(118);
				match(S_ASIGNAR);
				setState(119);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(false, id, entorno.NULL, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tipo_datoContext extends ParserRuleContext {
		public entorno.TipoDato tipo;
		public TerminalNode R_INT() { return getToken(sintactico.R_INT, 0); }
		public TerminalNode R_FLOAT() { return getToken(sintactico.R_FLOAT, 0); }
		public TerminalNode R_CHAR() { return getToken(sintactico.R_CHAR, 0); }
		public TerminalNode R_STRING() { return getToken(sintactico.R_STRING, 0); }
		public TerminalNode R_BOOL() { return getToken(sintactico.R_BOOL, 0); }
		public Tipo_datoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_dato; }
	}

	public final Tipo_datoContext tipo_dato() throws RecognitionException {
		Tipo_datoContext _localctx = new Tipo_datoContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_tipo_dato);
		try {
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(R_INT);
				 _localctx.tipo = entorno.INTEGER 
				}
				break;
			case R_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				match(R_FLOAT);
				 _localctx.tipo = entorno.FLOAT 
				}
				break;
			case R_CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(128);
				match(R_CHAR);
				 _localctx.tipo = entorno.CHAR 
				}
				break;
			case R_STRING:
				enterOuterAlt(_localctx, 4);
				{
				setState(130);
				match(R_STRING);
				 _localctx.tipo = entorno.STRING 
				}
				break;
			case R_BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(132);
				match(R_BOOL);
				 _localctx.tipo = entorno.BOOLEAN 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public LogicaContext logica;
		public RelacionalContext relacional;
		public AritmeticaContext aritmetica;
		public LogicaContext logica() {
			return getRuleContext(LogicaContext.class,0);
		}
		public RelacionalContext relacional() {
			return getRuleContext(RelacionalContext.class,0);
		}
		public AritmeticaContext aritmetica() {
			return getRuleContext(AritmeticaContext.class,0);
		}
		public ExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp; }
	}

	public final ExpContext exp() throws RecognitionException {
		ExpContext _localctx = new ExpContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_exp);
		try {
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				((ExpContext)_localctx).logica = logica(0);
				 _localctx.val = ((ExpContext)_localctx).logica.val 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(139);
				((ExpContext)_localctx).relacional = relacional(0);
				 _localctx.val = ((ExpContext)_localctx).relacional.val 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(142);
				((ExpContext)_localctx).aritmetica = aritmetica(0);
				 _localctx.val = ((ExpContext)_localctx).aritmetica.val 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LogicaContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public LogicaContext izq;
		public RelacionalContext relacional;
		public Token op;
		public LogicaContext der;
		public RelacionalContext relacional() {
			return getRuleContext(RelacionalContext.class,0);
		}
		public List<LogicaContext> logica() {
			return getRuleContexts(LogicaContext.class);
		}
		public LogicaContext logica(int i) {
			return getRuleContext(LogicaContext.class,i);
		}
		public TerminalNode S_AND() { return getToken(sintactico.S_AND, 0); }
		public TerminalNode S_OR() { return getToken(sintactico.S_OR, 0); }
		public LogicaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logica; }
	}

	public final LogicaContext logica() throws RecognitionException {
		return logica(0);
	}

	private LogicaContext logica(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		LogicaContext _localctx = new LogicaContext(_ctx, _parentState);
		LogicaContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_logica, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(148);
			((LogicaContext)_localctx).relacional = relacional(0);
			 _localctx.val = ((LogicaContext)_localctx).relacional.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(163);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(161);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new LogicaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_logica);
						setState(151);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(152);
						((LogicaContext)_localctx).op = match(S_AND);
						setState(153);
						((LogicaContext)_localctx).der = logica(4);

						          		_localctx.val = expresion.NewOperacion(((LogicaContext)_localctx).izq.val, ((LogicaContext)_localctx).der.val, (((LogicaContext)_localctx).op!=null?((LogicaContext)_localctx).op.getText():null), false)
						          	
						}
						break;
					case 2:
						{
						_localctx = new LogicaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_logica);
						setState(156);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(157);
						((LogicaContext)_localctx).op = match(S_OR);
						setState(158);
						((LogicaContext)_localctx).der = logica(3);

						          		_localctx.val = expresion.NewOperacion(((LogicaContext)_localctx).izq.val, ((LogicaContext)_localctx).der.val, (((LogicaContext)_localctx).op!=null?((LogicaContext)_localctx).op.getText():null), false)
						          	
						}
						break;
					}
					} 
				}
				setState(165);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class RelacionalContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public RelacionalContext izq;
		public AritmeticaContext aritmetica;
		public Token op;
		public RelacionalContext der;
		public AritmeticaContext aritmetica() {
			return getRuleContext(AritmeticaContext.class,0);
		}
		public List<RelacionalContext> relacional() {
			return getRuleContexts(RelacionalContext.class);
		}
		public RelacionalContext relacional(int i) {
			return getRuleContext(RelacionalContext.class,i);
		}
		public TerminalNode S_MENOR() { return getToken(sintactico.S_MENOR, 0); }
		public TerminalNode S_MAYOR() { return getToken(sintactico.S_MAYOR, 0); }
		public TerminalNode S_MENORI() { return getToken(sintactico.S_MENORI, 0); }
		public TerminalNode S_MAYORI() { return getToken(sintactico.S_MAYORI, 0); }
		public TerminalNode S_IGUAL() { return getToken(sintactico.S_IGUAL, 0); }
		public TerminalNode S_DIFERENTE() { return getToken(sintactico.S_DIFERENTE, 0); }
		public RelacionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relacional; }
	}

	public final RelacionalContext relacional() throws RecognitionException {
		return relacional(0);
	}

	private RelacionalContext relacional(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		RelacionalContext _localctx = new RelacionalContext(_ctx, _parentState);
		RelacionalContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_relacional, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(167);
			((RelacionalContext)_localctx).aritmetica = aritmetica(0);
			 _localctx.val = ((RelacionalContext)_localctx).aritmetica.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(177);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new RelacionalContext(_parentctx, _parentState);
					_localctx.izq = _prevctx;
					_localctx.izq = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_relacional);
					setState(170);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(171);
					((RelacionalContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << S_MAYOR) | (1L << S_MENOR) | (1L << S_MAYORI) | (1L << S_MENORI) | (1L << S_IGUAL) | (1L << S_DIFERENTE))) != 0)) ) {
						((RelacionalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(172);
					((RelacionalContext)_localctx).der = relacional(3);

					          		_localctx.val = expresion.NewOperacion(((RelacionalContext)_localctx).izq.val, ((RelacionalContext)_localctx).der.val, (((RelacionalContext)_localctx).op!=null?((RelacionalContext)_localctx).op.getText():null), false)
					          	
					}
					} 
				}
				setState(179);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AritmeticaContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public AritmeticaContext izq;
		public ExpContext exp;
		public Exp_valorContext exp_valor;
		public Token op;
		public AritmeticaContext der;
		public TerminalNode S_RESTA() { return getToken(sintactico.S_RESTA, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public Exp_valorContext exp_valor() {
			return getRuleContext(Exp_valorContext.class,0);
		}
		public TerminalNode S_APAR() { return getToken(sintactico.S_APAR, 0); }
		public TerminalNode S_CPAR() { return getToken(sintactico.S_CPAR, 0); }
		public List<AritmeticaContext> aritmetica() {
			return getRuleContexts(AritmeticaContext.class);
		}
		public AritmeticaContext aritmetica(int i) {
			return getRuleContext(AritmeticaContext.class,i);
		}
		public TerminalNode S_POR() { return getToken(sintactico.S_POR, 0); }
		public TerminalNode S_DIVISION() { return getToken(sintactico.S_DIVISION, 0); }
		public TerminalNode S_MODULO() { return getToken(sintactico.S_MODULO, 0); }
		public TerminalNode S_SUMA() { return getToken(sintactico.S_SUMA, 0); }
		public AritmeticaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aritmetica; }
	}

	public final AritmeticaContext aritmetica() throws RecognitionException {
		return aritmetica(0);
	}

	private AritmeticaContext aritmetica(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		AritmeticaContext _localctx = new AritmeticaContext(_ctx, _parentState);
		AritmeticaContext _prevctx = _localctx;
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_aritmetica, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case S_RESTA:
				{
				setState(181);
				match(S_RESTA);
				setState(182);
				((AritmeticaContext)_localctx).exp = exp();
				 _localctx.val = expresion.NewOperacion(((AritmeticaContext)_localctx).exp.val, nil, "-", true) 
				}
				break;
			case R_TRUE:
			case R_FALSE:
			case NUMERO:
			case DECIMAL:
			case CADENA:
			case ID:
				{
				setState(185);
				((AritmeticaContext)_localctx).exp_valor = exp_valor();
				 _localctx.val = ((AritmeticaContext)_localctx).exp_valor.val 
				}
				break;
			case S_APAR:
				{
				setState(188);
				match(S_APAR);
				setState(189);
				((AritmeticaContext)_localctx).exp = exp();
				setState(190);
				match(S_CPAR);
				 _localctx.val = ((AritmeticaContext)_localctx).exp.val 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(207);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(205);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
					case 1:
						{
						_localctx = new AritmeticaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_aritmetica);
						setState(195);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(196);
						((AritmeticaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << S_POR) | (1L << S_DIVISION) | (1L << S_MODULO))) != 0)) ) {
							((AritmeticaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(197);
						((AritmeticaContext)_localctx).der = aritmetica(5);

						          		_localctx.val = expresion.NewOperacion(((AritmeticaContext)_localctx).izq.val, ((AritmeticaContext)_localctx).der.val, (((AritmeticaContext)_localctx).op!=null?((AritmeticaContext)_localctx).op.getText():null), false) 
						          	
						}
						break;
					case 2:
						{
						_localctx = new AritmeticaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_aritmetica);
						setState(200);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(201);
						((AritmeticaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==S_SUMA || _la==S_RESTA) ) {
							((AritmeticaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(202);
						((AritmeticaContext)_localctx).der = aritmetica(4);

						          		_localctx.val = expresion.NewOperacion(((AritmeticaContext)_localctx).izq.val, ((AritmeticaContext)_localctx).der.val, (((AritmeticaContext)_localctx).op!=null?((AritmeticaContext)_localctx).op.getText():null), false) 
						          	
						}
						break;
					}
					} 
				}
				setState(209);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Exp_valorContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public PrimitivoContext primitivo;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public Exp_valorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp_valor; }
	}

	public final Exp_valorContext exp_valor() throws RecognitionException {
		Exp_valorContext _localctx = new Exp_valorContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_exp_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			((Exp_valorContext)_localctx).primitivo = primitivo();
			 _localctx.val = ((Exp_valorContext)_localctx).primitivo.val 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expresion val;
		public Token NUMERO;
		public Token DECIMAL;
		public Token CADENA;
		public Token ID;
		public TerminalNode NUMERO() { return getToken(sintactico.NUMERO, 0); }
		public TerminalNode DECIMAL() { return getToken(sintactico.DECIMAL, 0); }
		public TerminalNode CADENA() { return getToken(sintactico.CADENA, 0); }
		public TerminalNode R_TRUE() { return getToken(sintactico.R_TRUE, 0); }
		public TerminalNode R_FALSE() { return getToken(sintactico.R_FALSE, 0); }
		public TerminalNode ID() { return getToken(sintactico.ID, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_primitivo);
		try {
			setState(225);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				((PrimitivoContext)_localctx).NUMERO = match(NUMERO);

						val, err := strconv.Atoi((((PrimitivoContext)_localctx).NUMERO!=null?((PrimitivoContext)_localctx).NUMERO.getText():null))
						if err != nil {
							fmt.Println(err)
						}
						_localctx.val = expresion.NewPrimitivo(val, entorno.INTEGER)
					
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(215);
				((PrimitivoContext)_localctx).DECIMAL = match(DECIMAL);

						val, err := strconv.ParseFloat((((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getText():null), 64)
						if err != nil {
							fmt.Println(err)
						}
						_localctx.val = expresion.NewPrimitivo(val, entorno.FLOAT)
					
				}
				break;
			case CADENA:
				enterOuterAlt(_localctx, 3);
				{
				setState(217);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

						val := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1: len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)) - 1]
						_localctx.val = expresion.NewPrimitivo(val, entorno.STRING)
					
				}
				break;
			case R_TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(219);
				match(R_TRUE);

						_localctx.val = expresion.NewPrimitivo(true, entorno.BOOLEAN)
					
				}
				break;
			case R_FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(221);
				match(R_FALSE);

						_localctx.val = expresion.NewPrimitivo(false, entorno.BOOLEAN)
					
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(223);
				((PrimitivoContext)_localctx).ID = match(ID);

						_localctx.val = expresion.NewIdentificador((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null))
					
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return lista_exp_sempred((Lista_expContext)_localctx, predIndex);
		case 11:
			return logica_sempred((LogicaContext)_localctx, predIndex);
		case 12:
			return relacional_sempred((RelacionalContext)_localctx, predIndex);
		case 13:
			return aritmetica_sempred((AritmeticaContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean lista_exp_sempred(Lista_expContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean logica_sempred(LogicaContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 3);
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean relacional_sempred(RelacionalContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean aritmetica_sempred(AritmeticaContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 4);
		case 5:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3T\u00e6\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\3"+
		"\2\3\3\7\3\'\n\3\f\3\16\3*\13\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\6\6;\n\6\r\6\16\6<\3\6\3\6\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\5\7I\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\7\tZ\n\t\f\t\16\t]\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\5\n}\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\5\13\u0089\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u0094\n"+
		"\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\7\r\u00a4\n"+
		"\r\f\r\16\r\u00a7\13\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\7"+
		"\16\u00b2\n\16\f\16\16\16\u00b5\13\16\3\17\3\17\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u00c4\n\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\7\17\u00d0\n\17\f\17\16\17\u00d3\13\17"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\5\21\u00e4\n\21\3\21\2\6\20\30\32\34\22\2\4\6\b\n\f\16\20\22\24"+
		"\26\30\32\34\36 \2\5\3\2\20\25\3\2\r\17\3\2\13\f\2\u00ee\2\"\3\2\2\2\4"+
		"(\3\2\2\2\6-\3\2\2\2\b\60\3\2\2\2\n:\3\2\2\2\fH\3\2\2\2\16J\3\2\2\2\20"+
		"P\3\2\2\2\22|\3\2\2\2\24\u0088\3\2\2\2\26\u0093\3\2\2\2\30\u0095\3\2\2"+
		"\2\32\u00a8\3\2\2\2\34\u00c3\3\2\2\2\36\u00d4\3\2\2\2 \u00e3\3\2\2\2\""+
		"#\5\4\3\2#$\b\2\1\2$\3\3\2\2\2%\'\5\6\4\2&%\3\2\2\2\'*\3\2\2\2(&\3\2\2"+
		"\2()\3\2\2\2)+\3\2\2\2*(\3\2\2\2+,\b\3\1\2,\5\3\2\2\2-.\5\b\5\2./\b\4"+
		"\1\2/\7\3\2\2\2\60\61\7,\2\2\61\62\7-\2\2\62\63\7\31\2\2\63\64\7\32\2"+
		"\2\64\65\7\37\2\2\65\66\5\n\6\2\66\67\7 \2\2\678\b\5\1\28\t\3\2\2\29;"+
		"\5\f\7\2:9\3\2\2\2;<\3\2\2\2<:\3\2\2\2<=\3\2\2\2=>\3\2\2\2>?\b\6\1\2?"+
		"\13\3\2\2\2@A\5\16\b\2AB\7\5\2\2BC\b\7\1\2CI\3\2\2\2DE\5\22\n\2EF\7\5"+
		"\2\2FG\b\7\1\2GI\3\2\2\2H@\3\2\2\2HD\3\2\2\2I\r\3\2\2\2JK\7\63\2\2KL\7"+
		"\31\2\2LM\5\20\t\2MN\7\32\2\2NO\b\b\1\2O\17\3\2\2\2PQ\b\t\1\2QR\5\26\f"+
		"\2RS\b\t\1\2S[\3\2\2\2TU\f\4\2\2UV\7\4\2\2VW\5\26\f\2WX\b\t\1\2XZ\3\2"+
		"\2\2YT\3\2\2\2Z]\3\2\2\2[Y\3\2\2\2[\\\3\2\2\2\\\21\3\2\2\2][\3\2\2\2^"+
		"_\7$\2\2_`\7%\2\2`a\7R\2\2ab\7\33\2\2bc\5\24\13\2cd\7\6\2\2de\5\26\f\2"+
		"ef\b\n\1\2f}\3\2\2\2gh\7$\2\2hi\7%\2\2ij\7R\2\2jk\7\6\2\2kl\5\26\f\2l"+
		"m\b\n\1\2m}\3\2\2\2no\7$\2\2op\7R\2\2pq\7\33\2\2qr\5\24\13\2rs\7\6\2\2"+
		"st\5\26\f\2tu\b\n\1\2u}\3\2\2\2vw\7$\2\2wx\7R\2\2xy\7\6\2\2yz\5\26\f\2"+
		"z{\b\n\1\2{}\3\2\2\2|^\3\2\2\2|g\3\2\2\2|n\3\2\2\2|v\3\2\2\2}\23\3\2\2"+
		"\2~\177\7&\2\2\177\u0089\b\13\1\2\u0080\u0081\7\'\2\2\u0081\u0089\b\13"+
		"\1\2\u0082\u0083\7)\2\2\u0083\u0089\b\13\1\2\u0084\u0085\7*\2\2\u0085"+
		"\u0089\b\13\1\2\u0086\u0087\7(\2\2\u0087\u0089\b\13\1\2\u0088~\3\2\2\2"+
		"\u0088\u0080\3\2\2\2\u0088\u0082\3\2\2\2\u0088\u0084\3\2\2\2\u0088\u0086"+
		"\3\2\2\2\u0089\25\3\2\2\2\u008a\u008b\5\30\r\2\u008b\u008c\b\f\1\2\u008c"+
		"\u0094\3\2\2\2\u008d\u008e\5\32\16\2\u008e\u008f\b\f\1\2\u008f\u0094\3"+
		"\2\2\2\u0090\u0091\5\34\17\2\u0091\u0092\b\f\1\2\u0092\u0094\3\2\2\2\u0093"+
		"\u008a\3\2\2\2\u0093\u008d\3\2\2\2\u0093\u0090\3\2\2\2\u0094\27\3\2\2"+
		"\2\u0095\u0096\b\r\1\2\u0096\u0097\5\32\16\2\u0097\u0098\b\r\1\2\u0098"+
		"\u00a5\3\2\2\2\u0099\u009a\f\5\2\2\u009a\u009b\7\27\2\2\u009b\u009c\5"+
		"\30\r\6\u009c\u009d\b\r\1\2\u009d\u00a4\3\2\2\2\u009e\u009f\f\4\2\2\u009f"+
		"\u00a0\7\26\2\2\u00a0\u00a1\5\30\r\5\u00a1\u00a2\b\r\1\2\u00a2\u00a4\3"+
		"\2\2\2\u00a3\u0099\3\2\2\2\u00a3\u009e\3\2\2\2\u00a4\u00a7\3\2\2\2\u00a5"+
		"\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\31\3\2\2\2\u00a7\u00a5\3\2\2"+
		"\2\u00a8\u00a9\b\16\1\2\u00a9\u00aa\5\34\17\2\u00aa\u00ab\b\16\1\2\u00ab"+
		"\u00b3\3\2\2\2\u00ac\u00ad\f\4\2\2\u00ad\u00ae\t\2\2\2\u00ae\u00af\5\32"+
		"\16\5\u00af\u00b0\b\16\1\2\u00b0\u00b2\3\2\2\2\u00b1\u00ac\3\2\2\2\u00b2"+
		"\u00b5\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\33\3\2\2"+
		"\2\u00b5\u00b3\3\2\2\2\u00b6\u00b7\b\17\1\2\u00b7\u00b8\7\f\2\2\u00b8"+
		"\u00b9\5\26\f\2\u00b9\u00ba\b\17\1\2\u00ba\u00c4\3\2\2\2\u00bb\u00bc\5"+
		"\36\20\2\u00bc\u00bd\b\17\1\2\u00bd\u00c4\3\2\2\2\u00be\u00bf\7\31\2\2"+
		"\u00bf\u00c0\5\26\f\2\u00c0\u00c1\7\32\2\2\u00c1\u00c2\b\17\1\2\u00c2"+
		"\u00c4\3\2\2\2\u00c3\u00b6\3\2\2\2\u00c3\u00bb\3\2\2\2\u00c3\u00be\3\2"+
		"\2\2\u00c4\u00d1\3\2\2\2\u00c5\u00c6\f\6\2\2\u00c6\u00c7\t\3\2\2\u00c7"+
		"\u00c8\5\34\17\7\u00c8\u00c9\b\17\1\2\u00c9\u00d0\3\2\2\2\u00ca\u00cb"+
		"\f\5\2\2\u00cb\u00cc\t\4\2\2\u00cc\u00cd\5\34\17\6\u00cd\u00ce\b\17\1"+
		"\2\u00ce\u00d0\3\2\2\2\u00cf\u00c5\3\2\2\2\u00cf\u00ca\3\2\2\2\u00d0\u00d3"+
		"\3\2\2\2\u00d1\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\35\3\2\2\2\u00d3"+
		"\u00d1\3\2\2\2\u00d4\u00d5\5 \21\2\u00d5\u00d6\b\20\1\2\u00d6\37\3\2\2"+
		"\2\u00d7\u00d8\7N\2\2\u00d8\u00e4\b\21\1\2\u00d9\u00da\7O\2\2\u00da\u00e4"+
		"\b\21\1\2\u00db\u00dc\7Q\2\2\u00dc\u00e4\b\21\1\2\u00dd\u00de\7\"\2\2"+
		"\u00de\u00e4\b\21\1\2\u00df\u00e0\7#\2\2\u00e0\u00e4\b\21\1\2\u00e1\u00e2"+
		"\7R\2\2\u00e2\u00e4\b\21\1\2\u00e3\u00d7\3\2\2\2\u00e3\u00d9\3\2\2\2\u00e3"+
		"\u00db\3\2\2\2\u00e3\u00dd\3\2\2\2\u00e3\u00df\3\2\2\2\u00e3\u00e1\3\2"+
		"\2\2\u00e4!\3\2\2\2\20(<H[|\u0088\u0093\u00a3\u00a5\u00b3\u00c3\u00cf"+
		"\u00d1\u00e3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}