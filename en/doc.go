// Package en provides English version of ghts.
/*******************************************************************************
ghts
====

GHTS : GH Trading System

A software library for automatic trading system.

Library means that this is NOT a complete system,
but a collection of source code for (hopefully) useful 
developing a complete system.

You should develop 'buy&sell strategy' 
and risk management principle' by yourself.

Licensed under the term of GNU LGPL V3.
Refer to 'LICENSE' file, for the licensing detail

**************************************************************************

Copyright (C) 2014 UnHa Kim

This program is free software: 
you can redistribute it and/or modify it under the terms of the 
Version 3 of GNU Lesser General Public License
as published by the Free Software Foundation.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

See the GNU Lesser General Public License for more details.

You should have received a copy of the 
GNU Lesser General Public License (GNU LGPL) and 
GNU General Public License (GNU GPL) along with this program.  
If not, see <http://www.gnu.org/licenses/>.

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
But, mostly satisfied with balance of functionality and efficiency of Go.

Go have only 25 keywords at the of writing, 
which is much smaller number than Java and C++,
and even smaller than C and Python.

So it should be not too difficult for anyone to learn.

Give it a try.
*******************************************************************************/
package en
