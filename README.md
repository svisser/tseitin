go-tseitin
==========

Command-line tool written in Go implementing the Tseitin transformation.

Formula parser currently does not support Unicode or negation of formulas that
aren't literals. All formulas, including literals, need to be provided
with parentheses. 

Algorithm was published in:

- G.S. Tseitin: On the complexity of derivation in propositional calculus. In: Slisenko, A.O. (ed.) Structures in Constructive Mathematics and Mathematical Logic, Part II, Seminars in Mathematics (translated from Russian), pp. 115–125. Steklov Mathematical Institute (1968)
