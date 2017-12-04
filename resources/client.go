package resources 

import (
	"google.golang.org/grpc"
)

// TODO: implement methods below -- basically just call the corresponding grpc service method
// INTEGRATION 

func (i *Integration) NewClient(conn *grpc.ClientConn) {
	return NewIntegrationClient(conn)
}

func (i *Integration) Add(client *JobClient) {
	return nil
}

func (i *Integration) Get(client *JobClient) {
	return nil
}

func (i *Integration) List(query string, client *JobClient) {
	return nil
}

func (i *Integration) Remove(client *JobClient) {
	return nil
}

// REPOSITORY 

func (r *Repository) NewClient(conn *grpc.ClientConn) {
	return NewRepositoryClient(conn)
}

func (r *Repository) Add(client *JobClient) {
	return nil
}

func (r *Repository) Get(client *JobClient) {
	return nil
}

func (r *Repository) List(query string, client *JobClient) {
	return nil
}

func (r *Repository) Remove(client *JobClient) {
	return nil
}

// PIPELINE

func (p *Pipeline) NewClient(conn *grpc.ClientConn) {
	return NewPipelineClient(conn)
}

func (p *Pipeline) Run(client *JobClient) () {
	return nil
}

func (p *Pipeline) Get(client *JobClient) () {
	return nil
}

func (p *Pipeline) List(query string, client *JobClient) {
	return nil
}

func (p *Pipeline) Cancel(client *JobClient) {
	return nil
}

// JOB

func (j *Job) NewClient(conn *grpc.ClientConn) {
	return NewJobClient(conn)
}

func (j *Job) Get(client *JobClient) error {
	return nil
}

func (j *Job) List(query string, client *JobClient) error {
	return nil
}

// LOG

func (l *Log) NewClient(conn *grpc.ClientConn) {
	return NewLogClient(conn)
}

func (l *Log) Get(client *LogClient) error {
	return nil
}

// USER

func (u *User) NewClient(conn *grpc.ClientConn) {
	return NewUserClient(conn)
}

func (u *User) Add(client *UserClient) error {
	return nil
}