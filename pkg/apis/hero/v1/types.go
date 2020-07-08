/*
	types.go 文件的作用是定义一个 Hero 类型到底有哪些字段（比如，spec 字段里的内容）
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Hero 结构体描述一个 Hero 资源的字段
type Hero struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HeroSpec `json:"spec"`
}

// HeroSpec 描述一个具体的 Hero 资源的 spec 字段
type HeroSpec struct {
	Name string `json:"name"`
	Work string `json:"work"`
	Age  int16  `json:"age"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HeroList 是 Hero 的复数形式
type HeroList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Hero `json:"items"`
}
