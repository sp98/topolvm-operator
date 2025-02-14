/*
Copyright 2019 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8sutil

import (
	"context"
	"github.com/alauda/topolvm-operator/pkg/cluster"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestDeleteConfigMap(t *testing.T) {
	k8s := fake.NewSimpleClientset()

	cm := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-configmap",
			Namespace: "test-namespace",
			Labels:    map[string]string{cluster.LvmdConfigMapLabelKey: cluster.LvmdConfigMapLabelValue},
		},
		Data: map[string]string{
			"test": "data",
		},
	}

	ctx := context.TODO()

	_, err := k8s.CoreV1().ConfigMaps("test-namespace").Create(ctx, cm, metav1.CreateOptions{})
	assert.NoError(t, err)

	// There is no need to test all permutations, as the `DeleteResource` function is already
	// tested. Setting Wait=true and ErrorOnTimeout=true will cause both the delete and verify
	// functions to be exercised, and it will return error if either fail with an unexpected error.
	opts := &DeleteOptions{}
	opts.Wait = true
	opts.ErrorOnTimeout = true
	err = DeleteConfigMap(k8s, "test-configmap", "test-namespace", opts)
	assert.NoError(t, err)

	_, err = k8s.CoreV1().ConfigMaps("test-namespace").Get(ctx, "test-configmap", metav1.GetOptions{})
	assert.Error(t, err)
	assert.True(t, errors.IsNotFound(err))
}

func TestCreateReplaceableConfigmap(t *testing.T) {

	k8s := fake.NewSimpleClientset()
	cmOld := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-configmap",
			Namespace: "test-namespace",
		},
		Data: map[string]string{
			"test": "old-data",
		},
	}
	ctx := context.TODO()
	_, err := k8s.CoreV1().ConfigMaps("test-namespace").Create(ctx, cmOld, metav1.CreateOptions{})
	assert.NoError(t, err)

	cmNew := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-configmap",
			Namespace: "test-namespace",
		},
		Data: map[string]string{
			"test": "new-data",
		},
	}

	err = CreateReplaceableConfigmap(k8s, cmNew)
	assert.NoError(t, err)

	cm, err := k8s.CoreV1().ConfigMaps("test-namespace").Get(ctx, "test-configmap", metav1.GetOptions{})
	assert.NoError(t, err)

	assert.Equal(t, cmNew.Data["test"], cm.Data["test"])

}
