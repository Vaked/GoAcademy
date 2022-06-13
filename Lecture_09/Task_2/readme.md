Given the type:

type Action func() error

create a function called SafeExec, such that

func SafeExec(a Action) Action {
    // ...
}

returns a new function


By calling the new function, it will internally call the original one, and
wrap its error (if any), with the following message: "safe exec: %w"
Make it also handle possible panics using defer and recover.
Demonstrate in your main function, by passing a function without an
error, one with an error, and one that panics.

Hint: You will need to use a named returned value for the error, to be able to assign to it.