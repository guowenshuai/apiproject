/***
package controllers implement the controller for router

*/

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/guowenshuai/apiproject/models"
)

type JobController struct {
	beego.Controller
}

func (j *JobController) recoverHandler() {
	if err := recover(); err != nil {
		logs.Info("recover error msg: ", err)
		j.Data["json"] = map[string]string{"msg": "recover happened"}
		j.ServeJSON()
	}

}

// GetAll handles /job GET requests
// return all jobs
func (j *JobController) GetAll() {
	defer j.recoverHandler()
	logs.Info("jobs getAll ")
	_, jobs, err := models.GetAllJobs()
	if err != nil {
		j.Data["json"] = err.Error()
	}
	if len(jobs) == 0 {
		j.Data["json"] = map[string]string{
			"msg": "jobs is empty",
		}
	} else {
		j.Data["json"] = jobs
	}
	j.ServeJSON()
}

// Get handles /job/id GET requests
// return job record by job's field `Id`
func (j *JobController) Get() {
	defer j.recoverHandler()
	jobId := j.GetString(":jobId")
	job, err := models.JobById(jobId)
	if err != nil {
		beego.Error("find job by id: ", err.Error())
		j.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
	} else {
		j.Data["json"] = job
	}
	j.ServeJSON()
}

// ByBoss handles /job/ByBoss?boss=xxx GET requests
// return job record by job's field `boss`
func (j *JobController) ByBoss() {
	defer j.recoverHandler()
	boss := j.GetString(":boss")
	job, err := models.JobByBoss(boss)
	if err != nil {
		beego.Error("find job by boss: ", err.Error())
		j.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
	} else {
		j.Data["json"] = job
	}
	j.ServeJSON()
}

// Post handles /job POST requests
// insert a job record
func (j *JobController) Post() {
	defer j.recoverHandler()
	jobs := []models.Job{}
	json.Unmarshal(j.Ctx.Input.RequestBody, &jobs)

	insertId, err := models.InsertJobs(jobs)
	if err != nil {
		beego.Error("insert jobs error: ", err.Error())
		j.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
	} else {
		j.Data["json"] = map[string]int64{"success insert nums": insertId}
	}
	j.ServeJSON()
}

// Delete handles /job/id DELETE requests
// delete a job record by job's field Id
func (j *JobController) Delete() {
	jobId := j.GetString(":jobId")
	logs.Info("try to delete id: ", jobId)

	if job, err := models.JobDelete(jobId); err != nil {
		j.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
	} else {
		j.Data["json"] = job
	}
	j.ServeJSON()
}

// Put handles /job/id PUT requests
// update an exists job record
func (j *JobController) Put() {
	defer j.recoverHandler()
	jobId := j.GetString(":jobId")
	job := models.Job{}
	json.Unmarshal(j.Ctx.Input.RequestBody, &job)
	logs.Info("try to update id: ", jobId)

	if jb, err := models.JobUpdate(jobId, job); err != nil {
		logs.Info("update job failed: ", err.Error())
		j.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
	} else {
		j.Data["json"] = jb
	}
	j.ServeJSON()
}
