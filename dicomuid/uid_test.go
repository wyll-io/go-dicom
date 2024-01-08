package dicomuid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wyll-io/go-dicom/dicomuid"
)

func TestStandardUIDs(t *testing.T) {
	assert.Equal(t, dicomuid.PatientRootQRFind, "1.2.840.10008.5.1.4.1.2.1.1")
}

func TestLookupUID(t *testing.T) {
	u := dicomuid.MustLookup("1.2.840.10008.15.0.4.8")
	assert.Equal(t, u.Name, "dicomTransferCapability")
	assert.Equal(t, string(u.Type), "LDAP OID")
}
