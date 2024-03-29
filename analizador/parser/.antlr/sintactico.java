// Generated from c:\Users\Marco\Proyectos\Visual Code\goland\db_rust\analizador\parser\sintactico.g4 by ANTLR 4.8

	import "db_rust/analizador/ast"
	import "db_rust/analizador/ast/expresion"
	import "db_rust/analizador/ast/funcion"
	import "db_rust/analizador/ast/imprimir"
	import "db_rust/analizador/ast/variables"
	import "db_rust/analizador/ast/flujo"
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
		RULE_start = 0, RULE_procedimientos = 1, RULE_procedimiento = 2, RULE_principal = 3, 
		RULE_instrucciones = 4, RULE_instruccion = 5, RULE_imprimir = 6, RULE_lista_exp = 7, 
		RULE_declaracion = 8, RULE_tipo_dato = 9, RULE_asignacion = 10, RULE_sent_if = 11, 
		RULE_lista_elseif = 12, RULE_elseif = 13, RULE_sent_match = 14, RULE_casosMatch = 15, 
		RULE_casoMatch = 16, RULE_matchDefecto = 17, RULE_exp = 18, RULE_logica = 19, 
		RULE_relacional = 20, RULE_aritmetica = 21, RULE_exp_valor = 22, RULE_primitivo = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "procedimientos", "procedimiento", "principal", "instrucciones", 
			"instruccion", "imprimir", "lista_exp", "declaracion", "tipo_dato", "asignacion", 
			"sent_if", "lista_elseif", "elseif", "sent_match", "casosMatch", "casoMatch", 
			"matchDefecto", "exp", "logica", "relacional", "aritmetica", "exp_valor", 
			"primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "','", "';'", "'='", "'&'", "'|'", "'=>'", "'_'", "'+'", 
			"'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", 
			"'||'", "'&&'", "'->'", "'('", "')'", "':'", "'::'", "'['", "']'", "'{'", 
			"'}'", "'..'", "'true'", "'false'", "'let'", "'mut'", "'i64'", "'f64'", 
			"'bool'", "'char'", "'String'", "'&str'", "'fn'", "'main'", "'as'", "'to_owned'", 
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
		public ProcedimientosContext procedimientos;
		public ProcedimientosContext procedimientos() {
			return getRuleContext(ProcedimientosContext.class,0);
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
			setState(48);
			((StartContext)_localctx).procedimientos = procedimientos();

					_localctx.root = ast.NewAst(((StartContext)_localctx).procedimientos.lista)
				
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

	public static class ProcedimientosContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ProcedimientoContext procedimiento;
		public List<ProcedimientoContext> proc = new ArrayList<ProcedimientoContext>();
		public List<ProcedimientoContext> procedimiento() {
			return getRuleContexts(ProcedimientoContext.class);
		}
		public ProcedimientoContext procedimiento(int i) {
			return getRuleContext(ProcedimientoContext.class,i);
		}
		public ProcedimientosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_procedimientos; }
	}

	public final ProcedimientosContext procedimientos() throws RecognitionException {
		ProcedimientosContext _localctx = new ProcedimientosContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_procedimientos);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==R_FN) {
				{
				{
				setState(51);
				((ProcedimientosContext)_localctx).procedimiento = procedimiento();
				((ProcedimientosContext)_localctx).proc.add(((ProcedimientosContext)_localctx).procedimiento);
				}
				}
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

					LISTA := localctx.(*ProcedimientosContext).GetProc()
					for _, elemento := range LISTA {
						_localctx.lista.Add(elemento.GetInstr())
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

	public static class ProcedimientoContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public PrincipalContext principal;
		public PrincipalContext principal() {
			return getRuleContext(PrincipalContext.class,0);
		}
		public ProcedimientoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_procedimiento; }
	}

	public final ProcedimientoContext procedimiento() throws RecognitionException {
		ProcedimientoContext _localctx = new ProcedimientoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_procedimiento);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			((ProcedimientoContext)_localctx).principal = principal();
			 _localctx.instr = ((ProcedimientoContext)_localctx).principal.instr 
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

	public static class PrincipalContext extends ParserRuleContext {
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
		public PrincipalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_principal; }
	}

	public final PrincipalContext principal() throws RecognitionException {
		PrincipalContext _localctx = new PrincipalContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_principal);
		 params := arrayList.New() 
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(R_FN);
			setState(63);
			match(R_MAIN);
			setState(64);
			match(S_APAR);
			setState(65);
			match(S_CPAR);
			setState(66);
			match(S_ALLAV);
			setState(67);
			((PrincipalContext)_localctx).instrucciones = instrucciones();
			setState(68);
			match(S_CLLAV);

					_localctx.instr = funcion.NewFuncion(entorno.VOID, "main", params, ((PrincipalContext)_localctx).instrucciones.lista)
				
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
			setState(72); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(71);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).ins.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(74); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( ((((_la - 34)) & ~0x3f) == 0 && ((1L << (_la - 34)) & ((1L << (R_LET - 34)) | (1L << (R_PRINTLN - 34)) | (1L << (R_IF - 34)) | (1L << (R_MATCH - 34)) | (1L << (ID - 34)))) != 0) );

					LISTA := localctx.(*InstruccionesContext).GetIns()
					for _, elemento := range LISTA {
						_localctx.lista.Add(elemento.GetInstr())
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
		public AsignacionContext asignacion;
		public Sent_ifContext sent_if;
		public Sent_matchContext sent_match;
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public TerminalNode S_PTCOMA() { return getToken(sintactico.S_PTCOMA, 0); }
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public Sent_ifContext sent_if() {
			return getRuleContext(Sent_ifContext.class,0);
		}
		public Sent_matchContext sent_match() {
			return getRuleContext(Sent_matchContext.class,0);
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
			setState(98);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_PRINTLN:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				((InstruccionContext)_localctx).imprimir = imprimir();
				setState(79);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).imprimir.instr 
				}
				break;
			case R_LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(83);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).declaracion.instr 
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(86);
				((InstruccionContext)_localctx).asignacion = asignacion();
				setState(87);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).asignacion.instr 
				}
				break;
			case R_IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(90);
				((InstruccionContext)_localctx).sent_if = sent_if();
				setState(91);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).sent_if.instr 
				}
				break;
			case R_MATCH:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				((InstruccionContext)_localctx).sent_match = sent_match();
				setState(95);
				match(S_PTCOMA);
				 _localctx.instr = ((InstruccionContext)_localctx).sent_match.instr 
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
			setState(100);
			match(R_PRINTLN);
			setState(101);
			match(S_APAR);
			setState(102);
			((ImprimirContext)_localctx).lista_exp = lista_exp(0);
			setState(103);
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
			setState(107);
			((Lista_expContext)_localctx).exp = exp();
			 _localctx.lista.Add(((Lista_expContext)_localctx).exp.val) 
			}
			_ctx.stop = _input.LT(-1);
			setState(117);
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
					setState(110);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(111);
					match(S_COMA);
					setState(112);
					((Lista_expContext)_localctx).exp = exp();
					 
					          		((Lista_expContext)_localctx).LISTA.lista.Add(((Lista_expContext)_localctx).exp.val) 
					          		_localctx.lista = ((Lista_expContext)_localctx).LISTA.lista 
					          	
					}
					} 
				}
				setState(119);
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
			setState(150);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(R_LET);
				setState(121);
				match(R_MUT);
				setState(122);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(123);
				match(S_DOSPT);
				setState(124);
				((DeclaracionContext)_localctx).tipo_dato = tipo_dato();
				setState(125);
				match(S_ASIGNAR);
				setState(126);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(true, id, ((DeclaracionContext)_localctx).tipo_dato.tipo, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				match(R_LET);
				setState(130);
				match(R_MUT);
				setState(131);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(132);
				match(S_ASIGNAR);
				setState(133);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(true, id, entorno.NULL, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(136);
				match(R_LET);
				setState(137);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(138);
				match(S_DOSPT);
				setState(139);
				((DeclaracionContext)_localctx).tipo_dato = tipo_dato();
				setState(140);
				match(S_ASIGNAR);
				setState(141);
				((DeclaracionContext)_localctx).exp = exp();

						id := expresion.NewIdentificador((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null))
						_localctx.instr = variables.NewDeclaracion(false, id, ((DeclaracionContext)_localctx).tipo_dato.tipo, ((DeclaracionContext)_localctx).exp.val)
					
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(144);
				match(R_LET);
				setState(145);
				((DeclaracionContext)_localctx).ID = match(ID);
				setState(146);
				match(S_ASIGNAR);
				setState(147);
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
			setState(162);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				match(R_INT);
				 _localctx.tipo = entorno.INTEGER 
				}
				break;
			case R_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				match(R_FLOAT);
				 _localctx.tipo = entorno.FLOAT 
				}
				break;
			case R_CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(156);
				match(R_CHAR);
				 _localctx.tipo = entorno.CHAR 
				}
				break;
			case R_STRING:
				enterOuterAlt(_localctx, 4);
				{
				setState(158);
				match(R_STRING);
				 _localctx.tipo = entorno.STRING 
				}
				break;
			case R_BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(160);
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

	public static class AsignacionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token ID;
		public ExpContext exp;
		public TerminalNode ID() { return getToken(sintactico.ID, 0); }
		public TerminalNode S_ASIGNAR() { return getToken(sintactico.S_ASIGNAR, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(165);
			match(S_ASIGNAR);
			setState(166);
			((AsignacionContext)_localctx).exp = exp();

					id := expresion.NewIdentificador((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null))
					_localctx.instr = variables.NewAsignacion(id, ((AsignacionContext)_localctx).exp.val)
				
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

	public static class Sent_ifContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpContext exp;
		public InstruccionesContext s_then;
		public InstruccionesContext s_else;
		public Lista_elseifContext lista_elseif;
		public TerminalNode R_IF() { return getToken(sintactico.R_IF, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public List<TerminalNode> S_ALLAV() { return getTokens(sintactico.S_ALLAV); }
		public TerminalNode S_ALLAV(int i) {
			return getToken(sintactico.S_ALLAV, i);
		}
		public List<TerminalNode> S_CLLAV() { return getTokens(sintactico.S_CLLAV); }
		public TerminalNode S_CLLAV(int i) {
			return getToken(sintactico.S_CLLAV, i);
		}
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public TerminalNode R_ELSE() { return getToken(sintactico.R_ELSE, 0); }
		public Lista_elseifContext lista_elseif() {
			return getRuleContext(Lista_elseifContext.class,0);
		}
		public Sent_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sent_if; }
	}

	public final Sent_ifContext sent_if() throws RecognitionException {
		Sent_ifContext _localctx = new Sent_ifContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_sent_if);
		try {
			setState(199);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(R_IF);
				setState(170);
				((Sent_ifContext)_localctx).exp = exp();
				setState(171);
				match(S_ALLAV);
				setState(172);
				((Sent_ifContext)_localctx).s_then = instrucciones();
				setState(173);
				match(S_CLLAV);

						_localctx.instr = flujo.NewIf(((Sent_ifContext)_localctx).exp.val, ((Sent_ifContext)_localctx).s_then.lista, nil, nil)
					
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(176);
				match(R_IF);
				setState(177);
				((Sent_ifContext)_localctx).exp = exp();
				setState(178);
				match(S_ALLAV);
				setState(179);
				((Sent_ifContext)_localctx).s_then = instrucciones();
				setState(180);
				match(S_CLLAV);
				setState(181);
				match(R_ELSE);
				setState(182);
				match(S_ALLAV);
				setState(183);
				((Sent_ifContext)_localctx).s_else = instrucciones();
				setState(184);
				match(S_CLLAV);

						_localctx.instr = flujo.NewIf(((Sent_ifContext)_localctx).exp.val, ((Sent_ifContext)_localctx).s_then.lista, nil, ((Sent_ifContext)_localctx).s_else.lista)
					
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(187);
				match(R_IF);
				setState(188);
				((Sent_ifContext)_localctx).exp = exp();
				setState(189);
				match(S_ALLAV);
				setState(190);
				((Sent_ifContext)_localctx).s_then = instrucciones();
				setState(191);
				match(S_CLLAV);
				setState(192);
				((Sent_ifContext)_localctx).lista_elseif = lista_elseif();
				setState(193);
				match(R_ELSE);
				setState(194);
				match(S_ALLAV);
				setState(195);
				((Sent_ifContext)_localctx).s_else = instrucciones();
				setState(196);
				match(S_CLLAV);

						_localctx.instr = flujo.NewIf(((Sent_ifContext)_localctx).exp.val, ((Sent_ifContext)_localctx).s_then.lista, ((Sent_ifContext)_localctx).lista_elseif.lista, ((Sent_ifContext)_localctx).s_else.lista)
					
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

	public static class Lista_elseifContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ElseifContext elseif;
		public List<ElseifContext> ins = new ArrayList<ElseifContext>();
		public List<ElseifContext> elseif() {
			return getRuleContexts(ElseifContext.class);
		}
		public ElseifContext elseif(int i) {
			return getRuleContext(ElseifContext.class,i);
		}
		public Lista_elseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_elseif; }
	}

	public final Lista_elseifContext lista_elseif() throws RecognitionException {
		Lista_elseifContext _localctx = new Lista_elseifContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_lista_elseif);
		 _localctx.lista = arrayList.New() 
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(202); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(201);
					((Lista_elseifContext)_localctx).elseif = elseif();
					((Lista_elseifContext)_localctx).ins.add(((Lista_elseifContext)_localctx).elseif);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(204); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

					LISTA := localctx.(*Lista_elseifContext).GetIns()
					for _, elemento := range LISTA {
						_localctx.lista.Add(elemento.GetInstr())
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

	public static class ElseifContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpContext exp;
		public InstruccionesContext instrucciones;
		public TerminalNode R_ELSE() { return getToken(sintactico.R_ELSE, 0); }
		public TerminalNode R_IF() { return getToken(sintactico.R_IF, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode S_ALLAV() { return getToken(sintactico.S_ALLAV, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode S_CLLAV() { return getToken(sintactico.S_CLLAV, 0); }
		public ElseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseif; }
	}

	public final ElseifContext elseif() throws RecognitionException {
		ElseifContext _localctx = new ElseifContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_elseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(R_ELSE);
			setState(209);
			match(R_IF);
			setState(210);
			((ElseifContext)_localctx).exp = exp();
			setState(211);
			match(S_ALLAV);
			setState(212);
			((ElseifContext)_localctx).instrucciones = instrucciones();
			setState(213);
			match(S_CLLAV);

					_localctx.instr = flujo.NewIf(((ElseifContext)_localctx).exp.val, ((ElseifContext)_localctx).instrucciones.lista, nil, nil)
				
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

	public static class Sent_matchContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpContext exp;
		public CasosMatchContext casosMatch;
		public MatchDefectoContext matchDefecto;
		public TerminalNode R_MATCH() { return getToken(sintactico.R_MATCH, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode S_ALLAV() { return getToken(sintactico.S_ALLAV, 0); }
		public CasosMatchContext casosMatch() {
			return getRuleContext(CasosMatchContext.class,0);
		}
		public MatchDefectoContext matchDefecto() {
			return getRuleContext(MatchDefectoContext.class,0);
		}
		public TerminalNode S_CLLAV() { return getToken(sintactico.S_CLLAV, 0); }
		public Sent_matchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sent_match; }
	}

	public final Sent_matchContext sent_match() throws RecognitionException {
		Sent_matchContext _localctx = new Sent_matchContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_sent_match);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			match(R_MATCH);
			setState(217);
			((Sent_matchContext)_localctx).exp = exp();
			setState(218);
			match(S_ALLAV);
			setState(219);
			((Sent_matchContext)_localctx).casosMatch = casosMatch();
			setState(220);
			((Sent_matchContext)_localctx).matchDefecto = matchDefecto();
			setState(221);
			match(S_CLLAV);

					_localctx.instr = flujo.NewSentMatch(((Sent_matchContext)_localctx).exp.val, ((Sent_matchContext)_localctx).casosMatch.lista, ((Sent_matchContext)_localctx).matchDefecto.caso)
				
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

	public static class CasosMatchContext extends ParserRuleContext {
		public *arrayList.List lista;
		public CasoMatchContext casoMatch;
		public List<CasoMatchContext> caso = new ArrayList<CasoMatchContext>();
		public List<CasoMatchContext> casoMatch() {
			return getRuleContexts(CasoMatchContext.class);
		}
		public CasoMatchContext casoMatch(int i) {
			return getRuleContext(CasoMatchContext.class,i);
		}
		public CasosMatchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casosMatch; }
	}

	public final CasosMatchContext casosMatch() throws RecognitionException {
		CasosMatchContext _localctx = new CasosMatchContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_casosMatch);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(224);
				((CasosMatchContext)_localctx).casoMatch = casoMatch();
				((CasosMatchContext)_localctx).caso.add(((CasosMatchContext)_localctx).casoMatch);
				}
				}
				setState(227); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << S_RESTA) | (1L << S_APAR) | (1L << R_TRUE) | (1L << R_FALSE))) != 0) || ((((_la - 76)) & ~0x3f) == 0 && ((1L << (_la - 76)) & ((1L << (NUMERO - 76)) | (1L << (DECIMAL - 76)) | (1L << (CADENA - 76)) | (1L << (ID - 76)))) != 0) );

					LISTA := localctx.(*CasosMatchContext).GetCaso()
					for _, elemento := range LISTA {
						_localctx.lista.Add(elemento.GetCaso())
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

	public static class CasoMatchContext extends ParserRuleContext {
		public flujo.CaseMatch caso;
		public ExpContext exp;
		public InstruccionContext instruccion;
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode S_MATCH_RET() { return getToken(sintactico.S_MATCH_RET, 0); }
		public TerminalNode S_ALLAV() { return getToken(sintactico.S_ALLAV, 0); }
		public InstruccionContext instruccion() {
			return getRuleContext(InstruccionContext.class,0);
		}
		public TerminalNode S_CLLAV() { return getToken(sintactico.S_CLLAV, 0); }
		public TerminalNode S_COMA() { return getToken(sintactico.S_COMA, 0); }
		public CasoMatchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casoMatch; }
	}

	public final CasoMatchContext casoMatch() throws RecognitionException {
		CasoMatchContext _localctx = new CasoMatchContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_casoMatch);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			((CasoMatchContext)_localctx).exp = exp();
			setState(232);
			match(S_MATCH_RET);
			setState(233);
			match(S_ALLAV);
			setState(234);
			((CasoMatchContext)_localctx).instruccion = instruccion();
			setState(235);
			match(S_CLLAV);
			setState(236);
			match(S_COMA);

					_localctx.caso = flujo.NewCaseMatch(((CasoMatchContext)_localctx).exp.val, ((CasoMatchContext)_localctx).instruccion.instr)
				
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

	public static class MatchDefectoContext extends ParserRuleContext {
		public flujo.CaseMatch caso;
		public InstruccionContext instruccion;
		public TerminalNode S_MATCH_DEFAULT() { return getToken(sintactico.S_MATCH_DEFAULT, 0); }
		public TerminalNode S_MATCH_RET() { return getToken(sintactico.S_MATCH_RET, 0); }
		public TerminalNode S_ALLAV() { return getToken(sintactico.S_ALLAV, 0); }
		public InstruccionContext instruccion() {
			return getRuleContext(InstruccionContext.class,0);
		}
		public TerminalNode S_CLLAV() { return getToken(sintactico.S_CLLAV, 0); }
		public MatchDefectoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchDefecto; }
	}

	public final MatchDefectoContext matchDefecto() throws RecognitionException {
		MatchDefectoContext _localctx = new MatchDefectoContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_matchDefecto);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(S_MATCH_DEFAULT);
			setState(240);
			match(S_MATCH_RET);
			setState(241);
			match(S_ALLAV);
			setState(242);
			((MatchDefectoContext)_localctx).instruccion = instruccion();
			setState(243);
			match(S_CLLAV);

					_localctx.caso = flujo.NewCaseMatch(nil, ((MatchDefectoContext)_localctx).instruccion.instr)
				
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
		enterRule(_localctx, 36, RULE_exp);
		try {
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(246);
				((ExpContext)_localctx).logica = logica(0);
				 _localctx.val = ((ExpContext)_localctx).logica.val 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(249);
				((ExpContext)_localctx).relacional = relacional(0);
				 _localctx.val = ((ExpContext)_localctx).relacional.val 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(252);
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
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_logica, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(258);
			((LogicaContext)_localctx).relacional = relacional(0);
			 _localctx.val = ((LogicaContext)_localctx).relacional.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(273);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(271);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new LogicaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_logica);
						setState(261);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(262);
						((LogicaContext)_localctx).op = match(S_AND);
						setState(263);
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
						setState(266);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(267);
						((LogicaContext)_localctx).op = match(S_OR);
						setState(268);
						((LogicaContext)_localctx).der = logica(3);

						          		_localctx.val = expresion.NewOperacion(((LogicaContext)_localctx).izq.val, ((LogicaContext)_localctx).der.val, (((LogicaContext)_localctx).op!=null?((LogicaContext)_localctx).op.getText():null), false)
						          	
						}
						break;
					}
					} 
				}
				setState(275);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_relacional, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(277);
			((RelacionalContext)_localctx).aritmetica = aritmetica(0);
			 _localctx.val = ((RelacionalContext)_localctx).aritmetica.val 
			}
			_ctx.stop = _input.LT(-1);
			setState(287);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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
					setState(280);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(281);
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
					setState(282);
					((RelacionalContext)_localctx).der = relacional(3);

					          		_localctx.val = expresion.NewOperacion(((RelacionalContext)_localctx).izq.val, ((RelacionalContext)_localctx).der.val, (((RelacionalContext)_localctx).op!=null?((RelacionalContext)_localctx).op.getText():null), false)
					          	
					}
					} 
				}
				setState(289);
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
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_aritmetica, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case S_RESTA:
				{
				setState(291);
				match(S_RESTA);
				setState(292);
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
				setState(295);
				((AritmeticaContext)_localctx).exp_valor = exp_valor();
				 _localctx.val = ((AritmeticaContext)_localctx).exp_valor.val 
				}
				break;
			case S_APAR:
				{
				setState(298);
				match(S_APAR);
				setState(299);
				((AritmeticaContext)_localctx).exp = exp();
				setState(300);
				match(S_CPAR);
				 _localctx.val = ((AritmeticaContext)_localctx).exp.val 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(317);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(315);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new AritmeticaContext(_parentctx, _parentState);
						_localctx.izq = _prevctx;
						_localctx.izq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_aritmetica);
						setState(305);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(306);
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
						setState(307);
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
						setState(310);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(311);
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
						setState(312);
						((AritmeticaContext)_localctx).der = aritmetica(4);

						          		_localctx.val = expresion.NewOperacion(((AritmeticaContext)_localctx).izq.val, ((AritmeticaContext)_localctx).der.val, (((AritmeticaContext)_localctx).op!=null?((AritmeticaContext)_localctx).op.getText():null), false) 
						          	
						}
						break;
					}
					} 
				}
				setState(319);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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
		enterRule(_localctx, 44, RULE_exp_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(320);
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
		enterRule(_localctx, 46, RULE_primitivo);
		try {
			setState(335);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(323);
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
				setState(325);
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
				setState(327);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

						val := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1: len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)) - 1]
						_localctx.val = expresion.NewPrimitivo(val, entorno.STRING)
					
				}
				break;
			case R_TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(329);
				match(R_TRUE);

						_localctx.val = expresion.NewPrimitivo(true, entorno.BOOLEAN)
					
				}
				break;
			case R_FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(331);
				match(R_FALSE);

						_localctx.val = expresion.NewPrimitivo(false, entorno.BOOLEAN)
					
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(333);
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
		case 19:
			return logica_sempred((LogicaContext)_localctx, predIndex);
		case 20:
			return relacional_sempred((RelacionalContext)_localctx, predIndex);
		case 21:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3T\u0154\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\3\2\3\3\7\3\67\n\3\f\3\16\3:\13\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\6\6K\n\6\r\6\16\6L\3\6\3\6\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\5\7e\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\7\tv\n\t\f\t\16\ty\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\5\n\u0099\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5"+
		"\13\u00a5\n\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\5\r\u00ca\n\r\3\16\6\16\u00cd\n\16\r\16\16\16\u00ce\3\16"+
		"\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\21\6\21\u00e4\n\21\r\21\16\21\u00e5\3\21\3\21\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u0102\n\24\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u0112"+
		"\n\25\f\25\16\25\u0115\13\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\7\26\u0120\n\26\f\26\16\26\u0123\13\26\3\27\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0132\n\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u013e\n\27\f\27\16\27\u0141\13"+
		"\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\5\31\u0152\n\31\3\31\2\6\20(*,\32\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\2\5\3\2\20\25\3\2\r\17\3\2\13\f\2\u015b\2\62"+
		"\3\2\2\2\48\3\2\2\2\6=\3\2\2\2\b@\3\2\2\2\nJ\3\2\2\2\fd\3\2\2\2\16f\3"+
		"\2\2\2\20l\3\2\2\2\22\u0098\3\2\2\2\24\u00a4\3\2\2\2\26\u00a6\3\2\2\2"+
		"\30\u00c9\3\2\2\2\32\u00cc\3\2\2\2\34\u00d2\3\2\2\2\36\u00da\3\2\2\2 "+
		"\u00e3\3\2\2\2\"\u00e9\3\2\2\2$\u00f1\3\2\2\2&\u0101\3\2\2\2(\u0103\3"+
		"\2\2\2*\u0116\3\2\2\2,\u0131\3\2\2\2.\u0142\3\2\2\2\60\u0151\3\2\2\2\62"+
		"\63\5\4\3\2\63\64\b\2\1\2\64\3\3\2\2\2\65\67\5\6\4\2\66\65\3\2\2\2\67"+
		":\3\2\2\28\66\3\2\2\289\3\2\2\29;\3\2\2\2:8\3\2\2\2;<\b\3\1\2<\5\3\2\2"+
		"\2=>\5\b\5\2>?\b\4\1\2?\7\3\2\2\2@A\7,\2\2AB\7-\2\2BC\7\31\2\2CD\7\32"+
		"\2\2DE\7\37\2\2EF\5\n\6\2FG\7 \2\2GH\b\5\1\2H\t\3\2\2\2IK\5\f\7\2JI\3"+
		"\2\2\2KL\3\2\2\2LJ\3\2\2\2LM\3\2\2\2MN\3\2\2\2NO\b\6\1\2O\13\3\2\2\2P"+
		"Q\5\16\b\2QR\7\5\2\2RS\b\7\1\2Se\3\2\2\2TU\5\22\n\2UV\7\5\2\2VW\b\7\1"+
		"\2We\3\2\2\2XY\5\26\f\2YZ\7\5\2\2Z[\b\7\1\2[e\3\2\2\2\\]\5\30\r\2]^\7"+
		"\5\2\2^_\b\7\1\2_e\3\2\2\2`a\5\36\20\2ab\7\5\2\2bc\b\7\1\2ce\3\2\2\2d"+
		"P\3\2\2\2dT\3\2\2\2dX\3\2\2\2d\\\3\2\2\2d`\3\2\2\2e\r\3\2\2\2fg\7\63\2"+
		"\2gh\7\31\2\2hi\5\20\t\2ij\7\32\2\2jk\b\b\1\2k\17\3\2\2\2lm\b\t\1\2mn"+
		"\5&\24\2no\b\t\1\2ow\3\2\2\2pq\f\4\2\2qr\7\4\2\2rs\5&\24\2st\b\t\1\2t"+
		"v\3\2\2\2up\3\2\2\2vy\3\2\2\2wu\3\2\2\2wx\3\2\2\2x\21\3\2\2\2yw\3\2\2"+
		"\2z{\7$\2\2{|\7%\2\2|}\7R\2\2}~\7\33\2\2~\177\5\24\13\2\177\u0080\7\6"+
		"\2\2\u0080\u0081\5&\24\2\u0081\u0082\b\n\1\2\u0082\u0099\3\2\2\2\u0083"+
		"\u0084\7$\2\2\u0084\u0085\7%\2\2\u0085\u0086\7R\2\2\u0086\u0087\7\6\2"+
		"\2\u0087\u0088\5&\24\2\u0088\u0089\b\n\1\2\u0089\u0099\3\2\2\2\u008a\u008b"+
		"\7$\2\2\u008b\u008c\7R\2\2\u008c\u008d\7\33\2\2\u008d\u008e\5\24\13\2"+
		"\u008e\u008f\7\6\2\2\u008f\u0090\5&\24\2\u0090\u0091\b\n\1\2\u0091\u0099"+
		"\3\2\2\2\u0092\u0093\7$\2\2\u0093\u0094\7R\2\2\u0094\u0095\7\6\2\2\u0095"+
		"\u0096\5&\24\2\u0096\u0097\b\n\1\2\u0097\u0099\3\2\2\2\u0098z\3\2\2\2"+
		"\u0098\u0083\3\2\2\2\u0098\u008a\3\2\2\2\u0098\u0092\3\2\2\2\u0099\23"+
		"\3\2\2\2\u009a\u009b\7&\2\2\u009b\u00a5\b\13\1\2\u009c\u009d\7\'\2\2\u009d"+
		"\u00a5\b\13\1\2\u009e\u009f\7)\2\2\u009f\u00a5\b\13\1\2\u00a0\u00a1\7"+
		"*\2\2\u00a1\u00a5\b\13\1\2\u00a2\u00a3\7(\2\2\u00a3\u00a5\b\13\1\2\u00a4"+
		"\u009a\3\2\2\2\u00a4\u009c\3\2\2\2\u00a4\u009e\3\2\2\2\u00a4\u00a0\3\2"+
		"\2\2\u00a4\u00a2\3\2\2\2\u00a5\25\3\2\2\2\u00a6\u00a7\7R\2\2\u00a7\u00a8"+
		"\7\6\2\2\u00a8\u00a9\5&\24\2\u00a9\u00aa\b\f\1\2\u00aa\27\3\2\2\2\u00ab"+
		"\u00ac\7?\2\2\u00ac\u00ad\5&\24\2\u00ad\u00ae\7\37\2\2\u00ae\u00af\5\n"+
		"\6\2\u00af\u00b0\7 \2\2\u00b0\u00b1\b\r\1\2\u00b1\u00ca\3\2\2\2\u00b2"+
		"\u00b3\7?\2\2\u00b3\u00b4\5&\24\2\u00b4\u00b5\7\37\2\2\u00b5\u00b6\5\n"+
		"\6\2\u00b6\u00b7\7 \2\2\u00b7\u00b8\7@\2\2\u00b8\u00b9\7\37\2\2\u00b9"+
		"\u00ba\5\n\6\2\u00ba\u00bb\7 \2\2\u00bb\u00bc\b\r\1\2\u00bc\u00ca\3\2"+
		"\2\2\u00bd\u00be\7?\2\2\u00be\u00bf\5&\24\2\u00bf\u00c0\7\37\2\2\u00c0"+
		"\u00c1\5\n\6\2\u00c1\u00c2\7 \2\2\u00c2\u00c3\5\32\16\2\u00c3\u00c4\7"+
		"@\2\2\u00c4\u00c5\7\37\2\2\u00c5\u00c6\5\n\6\2\u00c6\u00c7\7 \2\2\u00c7"+
		"\u00c8\b\r\1\2\u00c8\u00ca\3\2\2\2\u00c9\u00ab\3\2\2\2\u00c9\u00b2\3\2"+
		"\2\2\u00c9\u00bd\3\2\2\2\u00ca\31\3\2\2\2\u00cb\u00cd\5\34\17\2\u00cc"+
		"\u00cb\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00cc\3\2\2\2\u00ce\u00cf\3\2"+
		"\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d1\b\16\1\2\u00d1\33\3\2\2\2\u00d2\u00d3"+
		"\7@\2\2\u00d3\u00d4\7?\2\2\u00d4\u00d5\5&\24\2\u00d5\u00d6\7\37\2\2\u00d6"+
		"\u00d7\5\n\6\2\u00d7\u00d8\7 \2\2\u00d8\u00d9\b\17\1\2\u00d9\35\3\2\2"+
		"\2\u00da\u00db\7A\2\2\u00db\u00dc\5&\24\2\u00dc\u00dd\7\37\2\2\u00dd\u00de"+
		"\5 \21\2\u00de\u00df\5$\23\2\u00df\u00e0\7 \2\2\u00e0\u00e1\b\20\1\2\u00e1"+
		"\37\3\2\2\2\u00e2\u00e4\5\"\22\2\u00e3\u00e2\3\2\2\2\u00e4\u00e5\3\2\2"+
		"\2\u00e5\u00e3\3\2\2\2\u00e5\u00e6\3\2\2\2\u00e6\u00e7\3\2\2\2\u00e7\u00e8"+
		"\b\21\1\2\u00e8!\3\2\2\2\u00e9\u00ea\5&\24\2\u00ea\u00eb\7\t\2\2\u00eb"+
		"\u00ec\7\37\2\2\u00ec\u00ed\5\f\7\2\u00ed\u00ee\7 \2\2\u00ee\u00ef\7\4"+
		"\2\2\u00ef\u00f0\b\22\1\2\u00f0#\3\2\2\2\u00f1\u00f2\7\n\2\2\u00f2\u00f3"+
		"\7\t\2\2\u00f3\u00f4\7\37\2\2\u00f4\u00f5\5\f\7\2\u00f5\u00f6\7 \2\2\u00f6"+
		"\u00f7\b\23\1\2\u00f7%\3\2\2\2\u00f8\u00f9\5(\25\2\u00f9\u00fa\b\24\1"+
		"\2\u00fa\u0102\3\2\2\2\u00fb\u00fc\5*\26\2\u00fc\u00fd\b\24\1\2\u00fd"+
		"\u0102\3\2\2\2\u00fe\u00ff\5,\27\2\u00ff\u0100\b\24\1\2\u0100\u0102\3"+
		"\2\2\2\u0101\u00f8\3\2\2\2\u0101\u00fb\3\2\2\2\u0101\u00fe\3\2\2\2\u0102"+
		"\'\3\2\2\2\u0103\u0104\b\25\1\2\u0104\u0105\5*\26\2\u0105\u0106\b\25\1"+
		"\2\u0106\u0113\3\2\2\2\u0107\u0108\f\5\2\2\u0108\u0109\7\27\2\2\u0109"+
		"\u010a\5(\25\6\u010a\u010b\b\25\1\2\u010b\u0112\3\2\2\2\u010c\u010d\f"+
		"\4\2\2\u010d\u010e\7\26\2\2\u010e\u010f\5(\25\5\u010f\u0110\b\25\1\2\u0110"+
		"\u0112\3\2\2\2\u0111\u0107\3\2\2\2\u0111\u010c\3\2\2\2\u0112\u0115\3\2"+
		"\2\2\u0113\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114)\3\2\2\2\u0115\u0113"+
		"\3\2\2\2\u0116\u0117\b\26\1\2\u0117\u0118\5,\27\2\u0118\u0119\b\26\1\2"+
		"\u0119\u0121\3\2\2\2\u011a\u011b\f\4\2\2\u011b\u011c\t\2\2\2\u011c\u011d"+
		"\5*\26\5\u011d\u011e\b\26\1\2\u011e\u0120\3\2\2\2\u011f\u011a\3\2\2\2"+
		"\u0120\u0123\3\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122+\3"+
		"\2\2\2\u0123\u0121\3\2\2\2\u0124\u0125\b\27\1\2\u0125\u0126\7\f\2\2\u0126"+
		"\u0127\5&\24\2\u0127\u0128\b\27\1\2\u0128\u0132\3\2\2\2\u0129\u012a\5"+
		".\30\2\u012a\u012b\b\27\1\2\u012b\u0132\3\2\2\2\u012c\u012d\7\31\2\2\u012d"+
		"\u012e\5&\24\2\u012e\u012f\7\32\2\2\u012f\u0130\b\27\1\2\u0130\u0132\3"+
		"\2\2\2\u0131\u0124\3\2\2\2\u0131\u0129\3\2\2\2\u0131\u012c\3\2\2\2\u0132"+
		"\u013f\3\2\2\2\u0133\u0134\f\6\2\2\u0134\u0135\t\3\2\2\u0135\u0136\5,"+
		"\27\7\u0136\u0137\b\27\1\2\u0137\u013e\3\2\2\2\u0138\u0139\f\5\2\2\u0139"+
		"\u013a\t\4\2\2\u013a\u013b\5,\27\6\u013b\u013c\b\27\1\2\u013c\u013e\3"+
		"\2\2\2\u013d\u0133\3\2\2\2\u013d\u0138\3\2\2\2\u013e\u0141\3\2\2\2\u013f"+
		"\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140-\3\2\2\2\u0141\u013f\3\2\2\2"+
		"\u0142\u0143\5\60\31\2\u0143\u0144\b\30\1\2\u0144/\3\2\2\2\u0145\u0146"+
		"\7N\2\2\u0146\u0152\b\31\1\2\u0147\u0148\7O\2\2\u0148\u0152\b\31\1\2\u0149"+
		"\u014a\7Q\2\2\u014a\u0152\b\31\1\2\u014b\u014c\7\"\2\2\u014c\u0152\b\31"+
		"\1\2\u014d\u014e\7#\2\2\u014e\u0152\b\31\1\2\u014f\u0150\7R\2\2\u0150"+
		"\u0152\b\31\1\2\u0151\u0145\3\2\2\2\u0151\u0147\3\2\2\2\u0151\u0149\3"+
		"\2\2\2\u0151\u014b\3\2\2\2\u0151\u014d\3\2\2\2\u0151\u014f\3\2\2\2\u0152"+
		"\61\3\2\2\2\238Ldw\u0098\u00a4\u00c9\u00ce\u00e5\u0101\u0111\u0113\u0121"+
		"\u0131\u013d\u013f\u0151";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}