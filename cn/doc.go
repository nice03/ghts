/*******************************************************************************
ghts
====

Chinese version of GHTS(GH Trading System).

LGPL V3 Licensed.


If you modified ghts code and distribute binary using ghts,
you should distribute source code of ghts itself and modification.
You can use this code in commercial software 
and keep your source code secret, except 'ghts' itself and 
its internal modification.

(Anyway ghts source code is already open. isn't it?)

**************************************************************************

Disclaimer of Warranties.

All of 'ghts', 'GHTS', 'GH Trading System' means same source code package.

Source code available in 'ghts' are provided "as is" 
  without warranty of any kind, 
  either expressed or implied 
  and such software is to be used at your own risk.

Authors(or developers) of 'ghts' disclaims to the fullest extent 
  authorized by law any and all other warranties, whether express or implied, 
  including, without limitation, any implied warranties of merchantability 
  or fitness for a particular purpose. 
  
The use of 'ghts' is done at your own discretion 
  and risk and with agreement 
  that you will be solely responsible for any damage or loss 
  to you and your computer.

You are solely responsible for adequate protection and backup of the data 
  and equipment used in any of the software related to 'ghts'. 
  and we will not be liable for any damages 
  that you may suffer in connection with downloading, installing, using, 
  modifying or distributing 'ghts'. 

No advice or information, whether oral or written, 
obtained by you from authors(or developers) of 'ghts' 
or from websites, 
or from source code,
or related documents
shall create any warranty for the software.
  
Without limitation of the foregoing, 
authors(or developers) of 'ghts' expressly does not warrant that:

1. the software will meet your requirements or expectations.
2. the software or the software content will be free of bugs, errors, 
     viruses or other defects.
3. any results, output, or data provided through or generated 
     by the software will be accurate, up-to-date, complete or reliable.
4. the software will be compatible with third party software.
5. any errors in the software will be corrected.

**************************************************************************

Basic structure

1. Price data module : Receive price data from broker,
						and distribute to other modules
2. Alpha module : Calculate probability or your trading strategy.
					Its objective is profit of interate rate + 'alpha'.
					(hence its name.)
3. Risk management and money management module :
					Limit potential damage or loss by watching portfolio.
4. Order management : send orders to broker.

Each modules are executed concurrently, not sequentially.
So, even when 'alpha module' is busy calculating trading strategy,
'price data module' have no problem receiving and distributing price data,
'risk management module' have no problem watching your portfolio and calculate
potential damage or loss or risk.ea
'order management module' have no problem sending orders.

In Go language, it is so easy to have goroutines execute concurrently.
And each module run in separate goroutine and that's all.

**************************************************************************

Naming convention.

First 1 or 2 character tells what type the data or function is.

It is easy to 'see'
1. mutable or immutable. (mutable : V, immutable : C)
2. interface or structure. (interface : I, C, V. structure : S, SC, SV)
3. normal function or constructor (normal function : F, constructor : NC, NV)

immutable data type is an easy way to prevent all the problem of 
sharing data in concurrent execution.

Scala language have compiler support immutable type declaration val,
and mutable type declaration var.
But, Go language doesn't support immutable type declaration.
But, wrap the value with struct,
and that struct does not have any public method which modified the value, 
then that value becomes immutable.
So, even withouth compiler support for immutable declaration,
custom immutable type is easy to create.

Anyway, Go compile speed is way more faster than Scala. isn't it?

I : I of Interface. Normal interface. 
C : C of Constant type. Interface for immutable data types.
V : V of Variable type. Interface for mutable data types.

F : F of Function. Normal function.
N : N of New instance. Constructor function.
NC : NC of New instance of Constant type. Constructor of immutable types.
NV : NV of New instance of Variable type. Constructor of mutable types.

S : S of Structure. struct of Go language.
SC : SC of Struct of Constant type. struct of immutable types.
SV : SC of Struct of Variable type. struct of mutable types.
G : G of Getter. Method which does not modify any of internal member fields.
	Usually have return value.
S : S of Setter. Method which set or modifies one or more of internal member fields.
	Often does not have any return value.

Go language compiler does not support constructor.
And new instance is created by new, which initialize all the member fields to zero value,
which often is not proper for some usage.

To have a constructor in Go languge, following pattern is used.
1. Concret structure is not public. 
	(First charact is in lower case which mean private scope in Go.)
2. Only constructor function can create new instance of structure.	
	And constructor function does all the necessary initialization.
3. All the struct is accessed only through related public interface.

So, using struct, function, interface, custom constructor is created.=

********************************************************************************

Go language

1. Who? : At first, Rob Pike and Ken Thomson.
			Rob Pike is legendary person of C language.
			De facto standard of C before ANSI C was 'K&R C' and 
			R of 'K&R C' means Rob Pike.
			Ken Thomson is famouse for co-father of UNIX.
			(Another UNIX creator Dennis Ritchie already passed away.)			

2. When? : late 2000s

3. When? : USA, Google Inc.

4. What? : A new programming language for high development efficiency, 
			hight execution speed, high multi-core CPU utilization.
			Often termed as golang.
			
5. Why? : Because sick of slow compilation and complexity of C++.
			With so many complex language functions, 
				there is no support for easy concurrency and parallelization 
				for multi-core CPU.
			Rob Pike and Ken Thomson decided that they need new programming 
			language for Google Inc.
			
6. How? : Based on C language grammar, removes crufts, and add
			lightweight object, (embedding not inheritance)
			automatic memory management, (garbage collection)
			easy concurrency and parallelization, (goroutine)
			have only necessary functions. 
			No more. only minimal. 
			But, fast compilation and high development efficiency as a results.
			Also, fast execution speed.

Some shorcomings
1. No or weak IDE(Integrated Development Environment) support.
	 
2. No intrinsic support for immutable data types.
	When mutable data is shared between concurrent execution units,
	ugly things happens.
	Immutable data type is easy answer.
	You can wrap data with structure and that structure does not have any public 
	method modifing internal member field, then the data become immutable.
	So a little inconvenient but solvable.
	
4. Type declaration is not familiar.
	Not '<type> <variable_name>', but '<variable_name> <type>'.
	For easy for compiler and resulting fast compilation speed.
	Need some time to get accustomed to.
	
5. Error is checked by if statement, and source code easily get dirty.
	Helper function helps to some degree.
	Or you can ignore any error return value, 
	and process error in function scope with 'defer & recover()'.
	Much like try, catch, finally in complete body of method scope.
	
6. Object type system is not based on inheritance but embedding.
	Need a little time to get familiar.
	
7. No method overloading.
	Each input parameter type have each function with different names or,
	One function accept all the type(with interface{}) or common interface,
	and type switch statement(switch <interface_name>.(type) {}) can handle
	different types differently.
	Personally prefer type switch statement with one function.
	
With all the shortcomings,
combined high development efficiency, high execution speed, easy concurrency
is really attractive to me.

When you do real-time trading,
can you accept sequencial processing?

No? then you need concurrency.
Thread? Event-driven? you are genius.
I am normal person and they are too difficult and complex for me.
I need something simple and easy.

Actor model(akka library in Java and Scala) and CSP(goroutine in Go) looks 
much simpler and easier.

Have experienced slow compilation of Java, 
Have experienced simple typing mistake create new variable in Python,
I choose Go langugage.

Sometime I miss some functionality in other language.
But, mostly satisfied with balance of functionality and efficiency in Go.

Give it a try.

Go have most small number of keyword in language grammar.
(Even compared to Python, which have so elegant and simple language grammar.)
So it should be not to difficult for anyone.
*******************************************************************************/
package cn
