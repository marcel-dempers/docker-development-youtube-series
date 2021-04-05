package main

import (
		"k8s.io/api/admission/v1beta1"
		"fmt"
		"encoding/json"
)

func Mutate(admissionReviewReq v1beta1.AdmissionReview) ( admissionReviewResponse v1beta1.AdmissionReview, err error) {

	admissionReviewResponse = v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			UID: admissionReviewReq.Request.UID,
			Allowed: true,
		},
	}

	if admissionReviewReq.Request.Operation == "CREATE" {
		fmt.Println("create event!")
		
		return MutateCreateJob(admissionReviewReq)
		// err = json.Unmarshal(admissionReviewReq.Request.Object.Raw, &pod)

		// if err != nil {
		// 	fmt.Errorf("could not unmarshal pod on admission request: %v", err)
		// 	return admissionReviewResponse, err
		// }
		
		// if ShouldMutuate(admissionReviewReq, pod){
		// 	return MutateCreatePod(pod, admissionReviewReq)
		// }
		
	}

	return admissionReviewResponse, nil
	
}

func MutateCreateJob(admissionReviewReq v1beta1.AdmissionReview) ( admissionReviewResponse v1beta1.AdmissionReview, err error) {
	
	admissionReviewResponse = v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			UID: admissionReviewReq.Request.UID,
			Allowed: true,
		},
	}

	out, err := json.Marshal(admissionReviewReq)
	if err != nil {
			panic (err)
	}

  fmt.Println(string(out))


	return admissionReviewResponse, nil
}
