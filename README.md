tseitin
========

Command-line tool written in Go implementing the Tseitin transformation.

This project was written to learn Go. It is not intended for use in production.

Formula parser currently does not support Unicode or negation of formulas that
aren't literals. All formulas, including literals, need to be provided
with parentheses. Lastly, literals cannot contain connectives in their name.
These restrictions can be fixed by using a better lexer / parser but that is
beyond the scope of this project.

Algorithm was published in:

- G.S. Tseitin: On the complexity of derivation in propositional calculus. In: Slisenko, A.O. (ed.) Structures in Constructive Mathematics and Mathematical Logic, Part II, Seminars in Mathematics (translated from Russian), pp. 115–125. Steklov Mathematical Institute (1968)
