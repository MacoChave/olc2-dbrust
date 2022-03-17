// Generated from c:\Users\Marco\Proyectos\Visual Code\goland\db_rust\analizador\parser\sintactico.g4 by ANTLR 4.8

	import "db_rust/analizador/ast"
	import "db_rust/analizador/ast/funcion"
	import "db_rust/analizador/ast/imprimir"
	import "db_rust/analizador/ast/expresion"
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
		RULE_exp = 8, RULE_logica = 9, RULE_relacional = 10, RULE_aritmetica = 11, 
		RULE_exp_valor = 12, RULE_primitivo = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funciones", "funcion", "funcMain", "instrucciones", "instruccion", 
			"imprimir", "lista_exp", "exp", "logica", "relacional", "aritmetica", 
			"exp_valor", "primitivo"
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
			setState(28);
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
			setState(34);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==R_FN) {
				{
				{
				setState(31);
				((FuncionesContext)_localctx).funcion = funcion();
				((FuncionesContext)_localctx).fun.add(((FuncionesContext)_localctx).funcion);
				}
				}
				setState(36);
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
			setState(39);
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
			setState(42);
			match(R_FN);
			setState(43);
			match(R_MAIN);
			setState(44);
			match(S_APAR);
			setState(45);
			match(S_CPAR);
			setState(46);
			match(S_ALLAV);
			setState(47);
			((FuncMainContext)_localctx).instrucciones = instrucciones();
			setState(48);
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
			setState(52); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(51);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).ins.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(54); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==R_PRINTLN );

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
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public TerminalNode S_PTCOMA() { return getToken(sintactico.S_PTCOMA, 0); }
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_instruccion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			((InstruccionContext)_localctx).imprimir = imprimir();
			setState(59);
			match(S_PTCOMA);
			 _localctx.instr = ((InstruccionContext)_localctx).imprimir.instr 
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
			setState(62);
			match(R_PRINTLN);
			setState(63);
			match(S_APAR);
			setState(64);
			((ImprimirContext)_localctx).lista_exp = lista_exp(0);
			setState(65);
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
			setState(69);
			((Lista_expContext)_localctx).exp = exp();
			 _localctx.lista.Add(((Lista_expContext)_localctx).exp.val) 
			}
			_ctx.stop = _input.LT(-1);
			setState(79);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
					setState(72);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(73);
					match(S_COMA);
					setState(74);
					((Lista_expContext)_localctx).exp = exp();
					 
					          		((Lista_expContext)_localctx).LISTA.lista.Add(((Lista_expContext)_localctx).exp.val) 
					          		_localctx.lista = ((Lista_expContext)_localctx).LISTA.lista 
					          	
					}
					} 
				}
				setState(81);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
		enterRule(_localctx, 16, RULE_exp);
		try {
			setState(91);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				((ExpContext)_localctx).logica = logica(0);
				 _localctx.val = ((ExpContext)_localctx).logica.val 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(85);
				((ExpContext)_localctx).relacional = relacional(0);
				 _localctx.val = ((ExpContext)_localctx).relacional.val 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(88);
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
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_logica, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(94);
			((LogicaContext)_localctx).relacional = relacional(0);
			 _localctx.val = ((LogicaContext)_localctx).relacional.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(109);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(107);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new LogicaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_logica);
						setState(97);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(98);
						((LogicaContext)_localctx).op = match(S_AND);
						setState(99);
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
						setState(102);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(103);
						((LogicaContext)_localctx).op = match(S_OR);
						setState(104);
						((LogicaContext)_localctx).der = logica(3);

						          		_localctx.val = expresion.NewOperacion(((LogicaContext)_localctx).izq.val, ((LogicaContext)_localctx).der.val, (((LogicaContext)_localctx).op!=null?((LogicaContext)_localctx).op.getText():null), false)
						          	
						}
						break;
					}
					} 
				}
				setState(111);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_relacional, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(113);
			((RelacionalContext)_localctx).aritmetica = aritmetica(0);
			 _localctx.val = ((RelacionalContext)_localctx).aritmetica.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(123);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
					setState(116);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(117);
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
					setState(118);
					((RelacionalContext)_localctx).der = relacional(3);

					          		_localctx.val = expresion.NewOperacion(((RelacionalContext)_localctx).izq.val, ((RelacionalContext)_localctx).der.val, (((RelacionalContext)_localctx).op!=null?((RelacionalContext)_localctx).op.getText():null), false)
					          	
					}
					} 
				}
				setState(125);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_aritmetica, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case S_RESTA:
				{
				setState(127);
				match(S_RESTA);
				setState(128);
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
				setState(131);
				((AritmeticaContext)_localctx).exp_valor = exp_valor();
				 _localctx.val = ((AritmeticaContext)_localctx).exp_valor.val 
				}
				break;
			case S_APAR:
				{
				setState(134);
				match(S_APAR);
				setState(135);
				((AritmeticaContext)_localctx).exp = exp();
				setState(136);
				match(S_CPAR);
				 _localctx.val = ((AritmeticaContext)_localctx).exp.val 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(153);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(151);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new AritmeticaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_aritmetica);
						setState(141);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(142);
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
						setState(143);
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
						setState(146);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(147);
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
						setState(148);
						((AritmeticaContext)_localctx).der = aritmetica(4);

						          		_localctx.val = expresion.NewOperacion(((AritmeticaContext)_localctx).izq.val, ((AritmeticaContext)_localctx).der.val, (((AritmeticaContext)_localctx).op!=null?((AritmeticaContext)_localctx).op.getText():null), false) 
						          	
						}
						break;
					}
					} 
				}
				setState(155);
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
		enterRule(_localctx, 24, RULE_exp_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
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
		enterRule(_localctx, 26, RULE_primitivo);
		try {
			setState(171);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(159);
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
				setState(161);
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
				setState(163);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

						val := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1: len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)) - 1]
						_localctx.val = expresion.NewPrimitivo(val, entorno.STRING)
					
				}
				break;
			case R_TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(165);
				match(R_TRUE);

						_localctx.val = expresion.NewPrimitivo(true, entorno.BOOLEAN)
					
				}
				break;
			case R_FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(167);
				match(R_FALSE);

						_localctx.val = expresion.NewPrimitivo(false, entorno.BOOLEAN)
					
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(169);
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
		case 9:
			return logica_sempred((LogicaContext)_localctx, predIndex);
		case 10:
			return relacional_sempred((RelacionalContext)_localctx, predIndex);
		case 11:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3T\u00b0\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\2\3\3\7\3#\n\3\f\3\16"+
		"\3&\13\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\6"+
		"\6\67\n\6\r\6\16\68\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\tP\n\t\f\t\16\tS\13\t\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\5\n^\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13n\n\13\f\13\16\13q\13\13\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f|\n\f\f\f\16\f\177\13\f\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u008e\n\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\7\r\u009a\n\r\f\r\16\r\u009d\13\r\3\16\3\16\3\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u00ae"+
		"\n\17\3\17\2\6\20\24\26\30\20\2\4\6\b\n\f\16\20\22\24\26\30\32\34\2\5"+
		"\3\2\20\25\3\2\r\17\3\2\13\f\2\u00b2\2\36\3\2\2\2\4$\3\2\2\2\6)\3\2\2"+
		"\2\b,\3\2\2\2\n\66\3\2\2\2\f<\3\2\2\2\16@\3\2\2\2\20F\3\2\2\2\22]\3\2"+
		"\2\2\24_\3\2\2\2\26r\3\2\2\2\30\u008d\3\2\2\2\32\u009e\3\2\2\2\34\u00ad"+
		"\3\2\2\2\36\37\5\4\3\2\37 \b\2\1\2 \3\3\2\2\2!#\5\6\4\2\"!\3\2\2\2#&\3"+
		"\2\2\2$\"\3\2\2\2$%\3\2\2\2%\'\3\2\2\2&$\3\2\2\2\'(\b\3\1\2(\5\3\2\2\2"+
		")*\5\b\5\2*+\b\4\1\2+\7\3\2\2\2,-\7,\2\2-.\7-\2\2./\7\31\2\2/\60\7\32"+
		"\2\2\60\61\7\37\2\2\61\62\5\n\6\2\62\63\7 \2\2\63\64\b\5\1\2\64\t\3\2"+
		"\2\2\65\67\5\f\7\2\66\65\3\2\2\2\678\3\2\2\28\66\3\2\2\289\3\2\2\29:\3"+
		"\2\2\2:;\b\6\1\2;\13\3\2\2\2<=\5\16\b\2=>\7\5\2\2>?\b\7\1\2?\r\3\2\2\2"+
		"@A\7\63\2\2AB\7\31\2\2BC\5\20\t\2CD\7\32\2\2DE\b\b\1\2E\17\3\2\2\2FG\b"+
		"\t\1\2GH\5\22\n\2HI\b\t\1\2IQ\3\2\2\2JK\f\4\2\2KL\7\4\2\2LM\5\22\n\2M"+
		"N\b\t\1\2NP\3\2\2\2OJ\3\2\2\2PS\3\2\2\2QO\3\2\2\2QR\3\2\2\2R\21\3\2\2"+
		"\2SQ\3\2\2\2TU\5\24\13\2UV\b\n\1\2V^\3\2\2\2WX\5\26\f\2XY\b\n\1\2Y^\3"+
		"\2\2\2Z[\5\30\r\2[\\\b\n\1\2\\^\3\2\2\2]T\3\2\2\2]W\3\2\2\2]Z\3\2\2\2"+
		"^\23\3\2\2\2_`\b\13\1\2`a\5\26\f\2ab\b\13\1\2bo\3\2\2\2cd\f\5\2\2de\7"+
		"\27\2\2ef\5\24\13\6fg\b\13\1\2gn\3\2\2\2hi\f\4\2\2ij\7\26\2\2jk\5\24\13"+
		"\5kl\b\13\1\2ln\3\2\2\2mc\3\2\2\2mh\3\2\2\2nq\3\2\2\2om\3\2\2\2op\3\2"+
		"\2\2p\25\3\2\2\2qo\3\2\2\2rs\b\f\1\2st\5\30\r\2tu\b\f\1\2u}\3\2\2\2vw"+
		"\f\4\2\2wx\t\2\2\2xy\5\26\f\5yz\b\f\1\2z|\3\2\2\2{v\3\2\2\2|\177\3\2\2"+
		"\2}{\3\2\2\2}~\3\2\2\2~\27\3\2\2\2\177}\3\2\2\2\u0080\u0081\b\r\1\2\u0081"+
		"\u0082\7\f\2\2\u0082\u0083\5\22\n\2\u0083\u0084\b\r\1\2\u0084\u008e\3"+
		"\2\2\2\u0085\u0086\5\32\16\2\u0086\u0087\b\r\1\2\u0087\u008e\3\2\2\2\u0088"+
		"\u0089\7\31\2\2\u0089\u008a\5\22\n\2\u008a\u008b\7\32\2\2\u008b\u008c"+
		"\b\r\1\2\u008c\u008e\3\2\2\2\u008d\u0080\3\2\2\2\u008d\u0085\3\2\2\2\u008d"+
		"\u0088\3\2\2\2\u008e\u009b\3\2\2\2\u008f\u0090\f\6\2\2\u0090\u0091\t\3"+
		"\2\2\u0091\u0092\5\30\r\7\u0092\u0093\b\r\1\2\u0093\u009a\3\2\2\2\u0094"+
		"\u0095\f\5\2\2\u0095\u0096\t\4\2\2\u0096\u0097\5\30\r\6\u0097\u0098\b"+
		"\r\1\2\u0098\u009a\3\2\2\2\u0099\u008f\3\2\2\2\u0099\u0094\3\2\2\2\u009a"+
		"\u009d\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\31\3\2\2"+
		"\2\u009d\u009b\3\2\2\2\u009e\u009f\5\34\17\2\u009f\u00a0\b\16\1\2\u00a0"+
		"\33\3\2\2\2\u00a1\u00a2\7N\2\2\u00a2\u00ae\b\17\1\2\u00a3\u00a4\7O\2\2"+
		"\u00a4\u00ae\b\17\1\2\u00a5\u00a6\7Q\2\2\u00a6\u00ae\b\17\1\2\u00a7\u00a8"+
		"\7\"\2\2\u00a8\u00ae\b\17\1\2\u00a9\u00aa\7#\2\2\u00aa\u00ae\b\17\1\2"+
		"\u00ab\u00ac\7R\2\2\u00ac\u00ae\b\17\1\2\u00ad\u00a1\3\2\2\2\u00ad\u00a3"+
		"\3\2\2\2\u00ad\u00a5\3\2\2\2\u00ad\u00a7\3\2\2\2\u00ad\u00a9\3\2\2\2\u00ad"+
		"\u00ab\3\2\2\2\u00ae\35\3\2\2\2\r$8Q]mo}\u008d\u0099\u009b\u00ad";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}