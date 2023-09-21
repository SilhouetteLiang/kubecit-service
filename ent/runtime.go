// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kubecit-service/ent/account"
	"kubecit-service/ent/category"
	"kubecit-service/ent/chapter"
	"kubecit-service/ent/course"
	"kubecit-service/ent/lesson"
	"kubecit-service/ent/orderinfos"
	"kubecit-service/ent/orders"
	"kubecit-service/ent/schema"
	"kubecit-service/ent/slider"
	"kubecit-service/ent/teacher"
	"kubecit-service/ent/wallet"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescOpenid is the schema descriptor for openid field.
	accountDescOpenid := accountFields[1].Descriptor()
	// account.OpenidValidator is a validator for the "openid" field. It is called by the builders before save.
	account.OpenidValidator = accountDescOpenid.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[2].Descriptor()
	// account.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = accountDescPassword.Validators[0].(func(string) error)
	// accountDescMethod is the schema descriptor for method field.
	accountDescMethod := accountFields[3].Descriptor()
	// account.MethodValidator is a validator for the "method" field. It is called by the builders before save.
	account.MethodValidator = accountDescMethod.Validators[0].(func(string) error)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	chapterFields := schema.Chapter{}.Fields()
	_ = chapterFields
	// chapterDescReleasedTime is the schema descriptor for released_time field.
	chapterDescReleasedTime := chapterFields[1].Descriptor()
	// chapter.DefaultReleasedTime holds the default value on creation for the released_time field.
	chapter.DefaultReleasedTime = chapterDescReleasedTime.Default.(func() time.Time)
	// chapter.UpdateDefaultReleasedTime holds the default value on update for the released_time field.
	chapter.UpdateDefaultReleasedTime = chapterDescReleasedTime.UpdateDefault.(func() time.Time)
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescUpdatedAt is the schema descriptor for updated_at field.
	courseDescUpdatedAt := courseFields[1].Descriptor()
	// course.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	course.DefaultUpdatedAt = courseDescUpdatedAt.Default.(func() time.Time)
	// course.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	course.UpdateDefaultUpdatedAt = courseDescUpdatedAt.UpdateDefault.(func() time.Time)
	// courseDescCreatedAt is the schema descriptor for created_at field.
	courseDescCreatedAt := courseFields[7].Descriptor()
	// course.DefaultCreatedAt holds the default value on creation for the created_at field.
	course.DefaultCreatedAt = courseDescCreatedAt.Default.(func() time.Time)
	// courseDescScore is the schema descriptor for score field.
	courseDescScore := courseFields[10].Descriptor()
	// course.DefaultScore holds the default value on creation for the score field.
	course.DefaultScore = courseDescScore.Default.(int32)
	// courseDescDuration is the schema descriptor for duration field.
	courseDescDuration := courseFields[11].Descriptor()
	// course.DefaultDuration holds the default value on creation for the duration field.
	course.DefaultDuration = courseDescDuration.Default.(int32)
	// courseDescPeople is the schema descriptor for people field.
	courseDescPeople := courseFields[12].Descriptor()
	// course.DefaultPeople holds the default value on creation for the people field.
	course.DefaultPeople = courseDescPeople.Default.(int32)
	lessonFields := schema.Lesson{}.Fields()
	_ = lessonFields
	// lessonDescReleasedTime is the schema descriptor for released_time field.
	lessonDescReleasedTime := lessonFields[1].Descriptor()
	// lesson.DefaultReleasedTime holds the default value on creation for the released_time field.
	lesson.DefaultReleasedTime = lessonDescReleasedTime.Default.(func() time.Time)
	// lesson.UpdateDefaultReleasedTime holds the default value on update for the released_time field.
	lesson.UpdateDefaultReleasedTime = lessonDescReleasedTime.UpdateDefault.(func() time.Time)
	// lessonDescIsFreePreview is the schema descriptor for is_free_preview field.
	lessonDescIsFreePreview := lessonFields[7].Descriptor()
	// lesson.DefaultIsFreePreview holds the default value on creation for the is_free_preview field.
	lesson.DefaultIsFreePreview = lessonDescIsFreePreview.Default.(int)
	orderinfosFields := schema.OrderInfos{}.Fields()
	_ = orderinfosFields
	// orderinfosDescCreateTime is the schema descriptor for create_time field.
	orderinfosDescCreateTime := orderinfosFields[5].Descriptor()
	// orderinfos.DefaultCreateTime holds the default value on creation for the create_time field.
	orderinfos.DefaultCreateTime = orderinfosDescCreateTime.Default.(func() time.Time)
	// orderinfosDescUpdateTime is the schema descriptor for update_time field.
	orderinfosDescUpdateTime := orderinfosFields[6].Descriptor()
	// orderinfos.DefaultUpdateTime holds the default value on creation for the update_time field.
	orderinfos.DefaultUpdateTime = orderinfosDescUpdateTime.Default.(func() time.Time)
	// orderinfos.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	orderinfos.UpdateDefaultUpdateTime = orderinfosDescUpdateTime.UpdateDefault.(func() time.Time)
	ordersFields := schema.Orders{}.Fields()
	_ = ordersFields
	// ordersDescPayType is the schema descriptor for pay_type field.
	ordersDescPayType := ordersFields[2].Descriptor()
	// orders.DefaultPayType holds the default value on creation for the pay_type field.
	orders.DefaultPayType = ordersDescPayType.Default.(int32)
	// ordersDescPayStatus is the schema descriptor for pay_status field.
	ordersDescPayStatus := ordersFields[3].Descriptor()
	// orders.DefaultPayStatus holds the default value on creation for the pay_status field.
	orders.DefaultPayStatus = ordersDescPayStatus.Default.(int32)
	// ordersDescCreateTime is the schema descriptor for create_time field.
	ordersDescCreateTime := ordersFields[7].Descriptor()
	// orders.DefaultCreateTime holds the default value on creation for the create_time field.
	orders.DefaultCreateTime = ordersDescCreateTime.Default.(func() time.Time)
	// ordersDescUpdateTime is the schema descriptor for update_time field.
	ordersDescUpdateTime := ordersFields[8].Descriptor()
	// orders.DefaultUpdateTime holds the default value on creation for the update_time field.
	orders.DefaultUpdateTime = ordersDescUpdateTime.Default.(func() time.Time)
	// orders.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	orders.UpdateDefaultUpdateTime = ordersDescUpdateTime.UpdateDefault.(func() time.Time)
	sliderFields := schema.Slider{}.Fields()
	_ = sliderFields
	// sliderDescTitle is the schema descriptor for title field.
	sliderDescTitle := sliderFields[0].Descriptor()
	// slider.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	slider.TitleValidator = sliderDescTitle.Validators[0].(func(string) error)
	// sliderDescContent is the schema descriptor for content field.
	sliderDescContent := sliderFields[1].Descriptor()
	// slider.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	slider.ContentValidator = sliderDescContent.Validators[0].(func(string) error)
	// sliderDescImageLink is the schema descriptor for image_link field.
	sliderDescImageLink := sliderFields[2].Descriptor()
	// slider.ImageLinkValidator is a validator for the "image_link" field. It is called by the builders before save.
	slider.ImageLinkValidator = sliderDescImageLink.Validators[0].(func(string) error)
	// sliderDescCreateAt is the schema descriptor for create_at field.
	sliderDescCreateAt := sliderFields[3].Descriptor()
	// slider.DefaultCreateAt holds the default value on creation for the create_at field.
	slider.DefaultCreateAt = sliderDescCreateAt.Default.(time.Time)
	// sliderDescUpdateAt is the schema descriptor for update_at field.
	sliderDescUpdateAt := sliderFields[4].Descriptor()
	// slider.DefaultUpdateAt holds the default value on creation for the update_at field.
	slider.DefaultUpdateAt = sliderDescUpdateAt.Default.(time.Time)
	// slider.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	slider.UpdateDefaultUpdateAt = sliderDescUpdateAt.UpdateDefault.(func() time.Time)
	// sliderDescIsValid is the schema descriptor for is_valid field.
	sliderDescIsValid := sliderFields[5].Descriptor()
	// slider.DefaultIsValid holds the default value on creation for the is_valid field.
	slider.DefaultIsValid = sliderDescIsValid.Default.(bool)
	teacherFields := schema.Teacher{}.Fields()
	_ = teacherFields
	// teacherDescCreateAt is the schema descriptor for create_at field.
	teacherDescCreateAt := teacherFields[7].Descriptor()
	// teacher.DefaultCreateAt holds the default value on creation for the create_at field.
	teacher.DefaultCreateAt = teacherDescCreateAt.Default.(func() time.Time)
	// teacherDescUpdateAt is the schema descriptor for update_at field.
	teacherDescUpdateAt := teacherFields[8].Descriptor()
	// teacher.DefaultUpdateAt holds the default value on creation for the update_at field.
	teacher.DefaultUpdateAt = teacherDescUpdateAt.Default.(func() time.Time)
	// teacher.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	teacher.UpdateDefaultUpdateAt = teacherDescUpdateAt.UpdateDefault.(func() time.Time)
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescGoldLeaf is the schema descriptor for gold_leaf field.
	walletDescGoldLeaf := walletFields[0].Descriptor()
	// wallet.DefaultGoldLeaf holds the default value on creation for the gold_leaf field.
	wallet.DefaultGoldLeaf = walletDescGoldLeaf.Default.(int32)
	// walletDescSilverLeaf is the schema descriptor for silver_leaf field.
	walletDescSilverLeaf := walletFields[1].Descriptor()
	// wallet.DefaultSilverLeaf holds the default value on creation for the silver_leaf field.
	wallet.DefaultSilverLeaf = walletDescSilverLeaf.Default.(int32)
	// walletDescFrozenGoldLeaf is the schema descriptor for frozen_gold_leaf field.
	walletDescFrozenGoldLeaf := walletFields[2].Descriptor()
	// wallet.DefaultFrozenGoldLeaf holds the default value on creation for the frozen_gold_leaf field.
	wallet.DefaultFrozenGoldLeaf = walletDescFrozenGoldLeaf.Default.(int32)
	// walletDescFrozenSilverLeaf is the schema descriptor for frozen_silver_leaf field.
	walletDescFrozenSilverLeaf := walletFields[3].Descriptor()
	// wallet.DefaultFrozenSilverLeaf holds the default value on creation for the frozen_silver_leaf field.
	wallet.DefaultFrozenSilverLeaf = walletDescFrozenSilverLeaf.Default.(int32)
	// walletDescCreatedAt is the schema descriptor for created_at field.
	walletDescCreatedAt := walletFields[5].Descriptor()
	// wallet.DefaultCreatedAt holds the default value on creation for the created_at field.
	wallet.DefaultCreatedAt = walletDescCreatedAt.Default.(func() time.Time)
	// walletDescUpdatedAt is the schema descriptor for updated_at field.
	walletDescUpdatedAt := walletFields[6].Descriptor()
	// wallet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	wallet.DefaultUpdatedAt = walletDescUpdatedAt.Default.(func() time.Time)
	// wallet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	wallet.UpdateDefaultUpdatedAt = walletDescUpdatedAt.UpdateDefault.(func() time.Time)
}
