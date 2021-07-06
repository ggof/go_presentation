package concurrent

import "golang.org/x/crypto/bcrypt"

type jobResult struct {
	Res []byte
	Err error
}

type job struct {
	password string
	result   chan jobResult
}

type JobSender = func(job) ([]byte, error)

func Start(number int) JobSender {
	jobs := make(chan job, number)

	for i := 0; i < number; i++ {
		go processWork(jobs)
	}

	return func(j job) ([]byte, error) {
		jobs <- j
		res := <-j.result
		return res.Res, res.Err
	}
}

func processWork(jobs <-chan job) {
	for j := range jobs {
		hash, err := bcrypt.GenerateFromPassword([]byte(j.password), bcrypt.DefaultCost)
		j.result <- jobResult{Res: hash, Err: err}
	}
}
