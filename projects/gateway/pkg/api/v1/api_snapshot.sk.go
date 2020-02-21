// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	VirtualServices VirtualServiceList
	RouteTables     RouteTableList
	Gateways        GatewayList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		VirtualServices: s.VirtualServices.Clone(),
		RouteTables:     s.RouteTables.Clone(),
		Gateways:        s.Gateways.Clone(),
	}
}

func (s ApiSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashVirtualServices(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRouteTables(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashGateways(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s ApiSnapshot) hashVirtualServices(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.VirtualServices.AsInterfaces()...)
}

func (s ApiSnapshot) hashRouteTables(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.RouteTables.AsInterfaces()...)
}

func (s ApiSnapshot) hashGateways(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Gateways.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	VirtualServicesHash, err := s.hashVirtualServices(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("virtualServices", VirtualServicesHash))
	RouteTablesHash, err := s.hashRouteTables(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("routeTables", RouteTablesHash))
	GatewaysHash, err := s.hashGateways(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("gateways", GatewaysHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type ApiSnapshotStringer struct {
	Version         uint64
	VirtualServices []string
	RouteTables     []string
	Gateways        []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  VirtualServices %v\n", len(ss.VirtualServices))
	for _, name := range ss.VirtualServices {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  RouteTables %v\n", len(ss.RouteTables))
	for _, name := range ss.RouteTables {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Gateways %v\n", len(ss.Gateways))
	for _, name := range ss.Gateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return ApiSnapshotStringer{
		Version:         snapshotHash,
		VirtualServices: s.VirtualServices.NamespacesDotNames(),
		RouteTables:     s.RouteTables.NamespacesDotNames(),
		Gateways:        s.Gateways.NamespacesDotNames(),
	}
}
