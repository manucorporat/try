package try_test

import (
	"github.com/manucorporat/try"
	"testing"
)

func Test_NormalFlow(T *testing.T) {
	called := false

	try.This(func() {
		called = true

	}).Catch(func(_ try.E) {
		T.Error("Catch must not be called")
	})

	// if try was not called
	if !called {
		T.Error("Try do not called")
	}
}

func Test_NormalFlowFinally(T *testing.T) {
	calledTry := false
	calledFinally := false

	try.This(func() {
		calledTry = true

	}).Finally(func() {
		calledFinally = true

	}).Catch(func(_ try.E) {
		T.Error("Catch must not be called")

	})

	// if try was not called
	if !calledTry {
		T.Error("Try do not called")
	}

	// if finally was not called
	if !calledFinally {
		T.Error("Finally do not called")
	}
}

func Test_CrashInTry(T *testing.T) {
	calledFinally := false
	calledCatch := false

	try.This(func() {
		panic("testing panic")

	}).Finally(func() {
		calledFinally = true

	}).Catch(func(e try.E) {
		calledCatch = true
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
	})

	// if catch was not called
	if !calledCatch {
		T.Error("Catch do not called")
	}

	// if finally was not called
	if !calledFinally {
		T.Error("Finally do not called")
	}
}

func Test_CrashInTry2(T *testing.T) {
	calledCatch := false

	try.This(func() {
		panic("testing panic")

	}).Catch(func(e try.E) {
		calledCatch = true
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
	})

	// if catch was not called
	if !calledCatch {
		T.Error("Catch do not called")
	}
}

func Test_CrashInCatch(T *testing.T) {
	calledFinally := false

	defer func() {
		err := recover()
		if err != "another panic" {
			T.Error("error is not 'another panic'")
		}
		// if finally was not called
		if !calledFinally {
			T.Error("Finally do not called")
		}
	}()
	try.This(func() {
		panic("testing panic")

	}).Finally(func() {
		calledFinally = true

	}).Catch(func(e try.E) {
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		panic("another panic")

	})
}

func Test_CrashInCatch2(T *testing.T) {
	defer func() {
		err := recover()
		if err != "another panic" {
			T.Error("error is not 'another panic'")
		}
	}()
	try.This(func() {
		panic("testing panic")

	}).Catch(func(e try.E) {
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		panic("another panic")
	})
}

func Test_CrashInThrow(T *testing.T) {
	calledFinally := false

	defer func() {
		err := recover()
		if err != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		// if finally was not called
		if !calledFinally {
			T.Error("Finally do not called")
		}
	}()

	try.This(func() {
		panic("testing panic")

	}).Finally(func() {
		calledFinally = true

	}).Catch(func(e try.E) {
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		try.Throw()

	})
}

func Test_CrashInThrow2(T *testing.T) {
	defer func() {
		err := recover()
		if err != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
	}()

	try.This(func() {
		panic("testing panic")

	}).Catch(func(e try.E) {
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		try.Throw()
	})
}

func Test_CrashInFinally1(T *testing.T) {
	calledTry := false

	defer func() {
		err := recover()
		if err != "finally panic" {
			T.Error("error is not 'finally panic'")
		}

		// if try was not called
		if !calledTry {
			T.Error("Try do not called")
		}
	}()

	try.This(func() {
		calledTry = true

	}).Finally(func() {
		panic("finally panic")

	}).Catch(func(e try.E) {
		T.Error("Catch must not be called")
	})
}

func Test_CrashInFinally2(T *testing.T) {

	defer func() {
		err := recover()
		if err != "finally panic" {
			T.Error("error is not 'finally panic'")
		}
	}()

	try.This(func() {
		panic("testing panic")

	}).Finally(func() {
		panic("finally panic")

	}).Catch(func(e try.E) {
		if e != "testing panic" {
			T.Error("error is not 'testing panic'")
		}
		panic("another panic")

	})
}
