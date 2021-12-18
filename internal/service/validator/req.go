/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 18:06 12月
 **/
package validator


import (
	v1 "Pheidippides/api/pheidiqueue/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

func QueueError(msg string) error {
	return errors.New(4000001, "queue err", msg)
}

func lackTopicError(msg string) error {
	return errors.New(4000002, "don't have topic", msg)
}

func Send(req *v1.Pheidippides) error {
	if req.GetTopic() == "" {
		return lackTopicError("send op must contains topic")
	}
	if req.GetContent() == "" {
		return errors.New(4000003, "don't have content", "send op must contains content")
	}
	return nil
}

func Get(req *v1.Pheidippides) error {
	if req.GetTopic() == "" {
		return lackTopicError("get op must contains topic")
	}
	if req.GetContent() != "" {
		return errors.New(4000004, "redundant content", "get op must't contains content")
	}
	return nil
}

