// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PollFactoryMetaData contains all meta data concerning the PollFactory contract.
var PollFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPoll\",\"inputs\":[{\"name\":\"question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"options\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"voterMerkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pollAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPoll\",\"inputs\":[{\"name\":\"pollId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPollId\",\"inputs\":[{\"name\":\"pollAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPolls\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pollCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pollIds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"polls\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PollCreated\",\"inputs\":[{\"name\":\"pollId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pollAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"question\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051611fbf380380611fbf83398101604081905261002e916100ae565b6001600160a01b0382166100555760405163baa3de5f60e01b815260040160405180910390fd5b6001600160a01b03811661007c57604051639589a27d60e01b815260040160405180910390fd5b6001600160a01b039182166080521660a0526100df565b80516001600160a01b03811681146100a9575f80fd5b919050565b5f80604083850312156100bf575f80fd5b6100c883610093565b91506100d660208401610093565b90509250929050565b60805160a051611eb161010e5f395f818161012b015261020b01525f81816101c001526101ea0152611eb15ff3fe608060405234801562000010575f80fd5b50600436106200009c575f3560e01c80639207891d116200006b5780639207891d146200014d578063ac2f00741462000156578063afea62311462000181578063c8cf9ab21462000198578063d6df096d14620001ba575f80fd5b80631a8cbcaa14620000a05780633ef87d5414620000e857806346e47f9a14620000fa5780637dc0d1d01462000125575b5f80fd5b620000cb620000b136600462000314565b5f908152600160205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020015b60405180910390f35b5f545b604051908152602001620000df565b620000eb6200010b3660046200032c565b6001600160a01b03165f9081526002602052604090205490565b620000cb7f000000000000000000000000000000000000000000000000000000000000000081565b620000eb5f5481565b620000cb6200016736600462000314565b60016020525f90815260409020546001600160a01b031681565b620000cb6200019236600462000416565b620001e2565b620000eb620001a93660046200032c565b60026020525f908152604090205481565b620000cb7f000000000000000000000000000000000000000000000000000000000000000081565b5f80858585857f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051620002389062000306565b620002499695949392919062000559565b604051809103905ff08015801562000263573d5f803e3d5ffd5b505f80549193508392509081806200027b8362000606565b909155505f81815260016020908152604080832080546001600160a01b0319166001600160a01b0389169081179091558084526002909252918290208390559051919250339183907f5b9a62314cee69a236b75c56d966d19a6903a71666f035bf55a6eec3292ade8890620002f4908c908b906200062b565b60405180910390a45050949350505050565b61182d806200064f83390190565b5f6020828403121562000325575f80fd5b5035919050565b5f602082840312156200033d575f80fd5b81356001600160a01b038116811462000354575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156200039b576200039b6200035b565b604052919050565b5f82601f830112620003b3575f80fd5b813567ffffffffffffffff811115620003d057620003d06200035b565b620003e5601f8201601f19166020016200036f565b818152846020838601011115620003fa575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f80608085870312156200042a575f80fd5b843567ffffffffffffffff8082111562000442575f80fd5b6200045088838901620003a3565b955060209150818701358181111562000467575f80fd5b8701601f8101891362000478575f80fd5b8035828111156200048d576200048d6200035b565b8060051b6200049e8582016200036f565b918252828101850191858101908c841115620004b8575f80fd5b86850192505b83831015620004f857823586811115620004d7575f8081fd5b620004e78e8983890101620003a3565b8352509186019190860190620004be565b999c999b50505050604088013597606001359695505050505050565b5f81518084525f5b818110156200053a576020818501810151868301820152016200051c565b505f602082860101526020601f19601f83011685010191505092915050565b60c081525f6200056d60c083018962000514565b6020838203818501528189518084528284019150828160051b850101838c015f5b83811015620005c057601f19878403018552620005ad83835162000514565b948601949250908501906001016200058e565b50508095505050505050856040830152846060830152620005ec60808301856001600160a01b03169052565b6001600160a01b03831660a0830152979650505050505050565b5f600182016200062457634e487b7160e01b5f52601160045260245ffd5b5060010190565b604081525f6200063f604083018562000514565b9050826020830152939250505056fe61016060405234801562000011575f80fd5b506040516200182d3803806200182d8339810160408190526200003491620004fd565b6002855110156200008c5760405162461bcd60e51b815260206004820152601b60248201527f4174206c656173742032206f7074696f6e73207265717569726564000000000060448201526064015b60405180910390fd5b5f8411620000dd5760405162461bcd60e51b815260206004820152601960248201527f4475726174696f6e206d75737420626520706f73697469766500000000000000604482015260640162000083565b826200012c5760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964204d65726b6c6520726f6f7400000000000000000000000000604482015260640162000083565b6001600160a01b038216620001775760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b2103b32b934b334b2b960811b604482015260640162000083565b6001600160a01b038116620001c05760405162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f7261636c6560901b604482015260640162000083565b336080526001620001d28782620006b5565b508451620001e8906002906020880190620002f8565b5060c08490526101008390526001600160a01b03808316610120528116610140524260a08190526200021c9085906200077d565b60e0525f805460ff1916905584516001600160401b0381111562000244576200024462000412565b6040519080825280602002602001820160405280156200026e578160200160208202803683370190505b508051620002859160049160209091019062000353565b506101405160e051604051637b13423360e01b815230600482015260248101919091526001600160a01b0390911690637b134233906044015f604051808303815f87803b158015620002d5575f80fd5b505af1158015620002e8573d5f803e3d5ffd5b50505050505050505050620007a3565b828054828255905f5260205f2090810192821562000341579160200282015b82811115620003415782518290620003309082620006b5565b509160200191906001019062000317565b506200034f9291506200039d565b5090565b828054828255905f5260205f209081019282156200038f579160200282015b828111156200038f57825182559160200191906001019062000372565b506200034f929150620003bd565b808211156200034f575f620003b38282620003d3565b506001016200039d565b5b808211156200034f575f8155600101620003be565b508054620003e19062000629565b5f825580601f10620003f1575050565b601f0160209004905f5260205f20908101906200040f9190620003bd565b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171562000451576200045162000412565b604052919050565b5f82601f83011262000469575f80fd5b81516001600160401b0381111562000485576200048562000412565b60206200049b601f8301601f1916820162000426565b8281528582848701011115620004af575f80fd5b5f5b83811015620004ce578581018301518282018401528201620004b1565b505f928101909101919091529392505050565b80516001600160a01b0381168114620004f8575f80fd5b919050565b5f805f805f8060c0878903121562000513575f80fd5b86516001600160401b03808211156200052a575f80fd5b620005388a838b0162000459565b975060208901519150808211156200054e575f80fd5b818901915089601f83011262000562575f80fd5b81518181111562000577576200057762000412565b8060051b620005896020820162000426565b9182526020818501810192908101908d841115620005a5575f80fd5b6020860192505b83831015620005e857825185811115620005c4575f80fd5b620005d58f6020838a010162000459565b83525060209283019290910190620005ac565b809a5050505050505060408701519350606087015192506200060d60808801620004e1565b91506200061d60a08801620004e1565b90509295509295509295565b600181811c908216806200063e57607f821691505b6020821081036200065d57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620006b0575f81815260208120601f850160051c810160208610156200068b5750805b601f850160051c820191505b81811015620006ac5782815560010162000697565b5050505b505050565b81516001600160401b03811115620006d157620006d162000412565b620006e981620006e2845462000629565b8462000663565b602080601f8311600181146200071f575f8415620007075750858301515b5f19600386901b1c1916600185901b178555620006ac565b5f85815260208120601f198616915b828110156200074f578886015182559484019460019091019084016200072e565b50858210156200076d57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200079d57634e487b7160e01b5f52601160045260245ffd5b92915050565b60805160a05160c05160e0516101005161012051610140516110236200080a5f395f818161026f0152610aa901525f81816102e9015261094501525f81816101c70152610c3c01525f6101ee01525f61017401525f6102c201525f61013001526110235ff3fe608060405234801561000f575f80fd5b5060043610610127575f3560e01c8063713916bd116100a9578063d6df096d1161006e578063d6df096d146102e4578063e8fcf7231461030b578063ed8c2aed14610359578063f373856914610361578063fa1026dd1461036a575f80fd5b8063713916bd146102575780637dc0d1d01461026a578063a254576114610291578063c19d93fb146102a4578063cf09e0d0146102bd575f80fd5b80633197cbb6116100ef5780633197cbb6146101e95780633fad9ae014610210578063410673e5146102255780634717f97c1461023a57806369edbcda14610242575f80fd5b806302d05d3f1461012b5780630fb5a6b41461016f5780631069143a146101a45780631d3231d4146101b95780632a69fb46146101c2575b5f80fd5b6101527f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101967f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610166565b6101ac6103f6565b6040516101669190610ca8565b61019660055481565b6101967f000000000000000000000000000000000000000000000000000000000000000081565b6101967f000000000000000000000000000000000000000000000000000000000000000081565b6102186104ca565b6040516101669190610d08565b61022d610556565b6040516101669190610d5a565b61022d610629565b610255610250366004610d6c565b6106b3565b005b610255610265366004610d8c565b610863565b6101527f000000000000000000000000000000000000000000000000000000000000000081565b61021861029f366004610e55565b610a77565b5f546102b09060ff1681565b6040516101669190610e80565b6101967f000000000000000000000000000000000000000000000000000000000000000081565b6101527f000000000000000000000000000000000000000000000000000000000000000081565b61033c610319366004610ea6565b60036020525f908152604090208054600182015460029092015490919060ff1683565b604080519384526020840192909252151590820152606001610166565b610255610a9e565b61019660065481565b6103d2610378366004610ea6565b60408051606080820183525f80835260208084018290529284018190526001600160a01b03949094168452600382529282902082519384018352805484526001810154918401919091526002015460ff1615159082015290565b60408051825181526020808401519082015291810151151590820152606001610166565b60606002805480602002602001604051908101604052809291908181526020015f905b828210156104c1578382905f5260205f2001805461043690610ecc565b80601f016020809104026020016040519081016040528092919081815260200182805461046290610ecc565b80156104ad5780601f10610484576101008083540402835291602001916104ad565b820191905f5260205f20905b81548152906001019060200180831161049057829003601f168201915b505050505081526020019060010190610419565b50505050905090565b600180546104d790610ecc565b80601f016020809104026020016040519081016040528092919081815260200182805461050390610ecc565b801561054e5780601f106105255761010080835404028352916020019161054e565b820191905f5260205f20905b81548152906001019060200180831161053157829003601f168201915b505050505081565b606060015f5460ff16600281111561057057610570610e6c565b1461058e576040516387da8e5f60e01b815260040160405180910390fd5b5f805460ff191660021790556040517f5408f1d0b104f82c19a2174a2056a36345c26e44e0e7df271f5aff18e9a6ce44906105cd906004904290610f04565b60405180910390a1600480548060200260200160405190810160405280929190818152602001828054801561061f57602002820191905f5260205f20905b81548152602001906001019080831161060b575b5050505050905090565b606060025f5460ff16600281111561064357610643610e6c565b1461066157604051638065a29d60e01b815260040160405180910390fd5b600480548060200260200160405190810160405280929190818152602001828054801561061f57602002820191905f5260205f209081548152602001906001019080831161060b575050505050905090565b60015f5460ff1660028111156106cb576106cb610e6c565b146106e9576040516387da8e5f60e01b815260040160405180910390fd5b335f908152600360205260409020805461071657604051635b07c98960e01b815260040160405180910390fd5b600281015460ff161561073c5760405163a89ac15160e01b815260040160405180910390fd5b600254831061075e57604051639c45400160e01b815260040160405180910390fd5b5f83833360405160200161079793929190928352602083019190915260601b6bffffffffffffffffffffffff1916604082015260540190565b604051602081830303815290604052805190602001209050815f015481146107d257604051639ea6d12760e01b815260040160405180910390fd5b60028201805460ff1916600117905560048054859081106107f5576107f5610f50565b5f918252602082200180549161080a83610f64565b909155505060068054905f61081e83610f64565b90915550506040805185815242602082015233917ff65a04be847d83385f9a7abcf32cfb35055023a4affcc7eaa319160746f44528910160405180910390a250505050565b5f805460ff16600281111561087a5761087a610e6c565b14610898576040516327fd0eed60e11b815260040160405180910390fd5b335f90815260036020526040902054156108c557604051637c9a1cf960e01b815260040160405180910390fd5b6108d0338383610b5d565b6108ed5760405163582f497d60e11b815260040160405180910390fd5b6040805160018082528183019092525f9160208083019080368337019050509050855f1c815f8151811061092357610923610f50565b6020908102919091010152604051631e8e1e1360e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631e8e1e139061097e90889088908690600401610f88565b602060405180830381865afa158015610999573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109bd9190610fce565b6109da576040516309bde33960e01b815260040160405180910390fd5b604080516060810182528781524260208083019182525f83850181815233825260039092529384209251835590516001830155516002909101805460ff19169115159190911790556005805491610a3083610f64565b90915550506040805187815242602082015233917f06fdcf30f0bb2c4ceab314db052ad51b198e3ff93ce4121183516fa1ae84fbc8910160405180910390a2505050505050565b60028181548110610a86575f80fd5b905f5260205f20015f9150905080546104d790610ecc565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ae75760405163db8d1fb760e01b815260040160405180910390fd5b5f805460ff166002811115610afe57610afe610e6c565b14610b1c57604051639b22ddeb60e01b815260040160405180910390fd5b5f805460ff191660011790556040514281527f703c67c157a1afb7d386ce5212ee88bf0a690bca998060e9595ebdfd747100ce9060200160405180910390a1565b6040516bffffffffffffffffffffffff19606085901b1660208201525f90819060340160408051601f1981840301815291905280516020909101209050805f5b84811015610c39575f868683818110610bb857610bb8610f50565b905060200201359050808311610bf9576040805160208101859052908101829052606001604051602081830303815290604052805190602001209250610c26565b60408051602081018390529081018490526060016040516020818303038152906040528051906020012092505b5080610c3181610f64565b915050610b9d565b507f00000000000000000000000000000000000000000000000000000000000000001495945050505050565b5f81518084525f5b81811015610c8957602081850181015186830182015201610c6d565b505f602082860101526020601f19601f83011685010191505092915050565b5f602080830181845280855180835260408601915060408160051b87010192508387015f5b82811015610cfb57603f19888603018452610ce9858351610c65565b94509285019290850190600101610ccd565b5092979650505050505050565b602081525f610d1a6020830184610c65565b9392505050565b5f8151808452602080850194508084015f5b83811015610d4f57815187529582019590820190600101610d33565b509495945050505050565b602081525f610d1a6020830184610d21565b5f8060408385031215610d7d575f80fd5b50508035926020909101359150565b5f805f805f60608688031215610da0575f80fd5b85359450602086013567ffffffffffffffff80821115610dbe575f80fd5b818801915088601f830112610dd1575f80fd5b813581811115610ddf575f80fd5b896020828501011115610df0575f80fd5b602083019650809550506040880135915080821115610e0d575f80fd5b818801915088601f830112610e20575f80fd5b813581811115610e2e575f80fd5b8960208260051b8501011115610e42575f80fd5b9699959850939650602001949392505050565b5f60208284031215610e65575f80fd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b6020810160038310610ea057634e487b7160e01b5f52602160045260245ffd5b91905290565b5f60208284031215610eb6575f80fd5b81356001600160a01b0381168114610d1a575f80fd5b600181811c90821680610ee057607f821691505b602082108103610efe57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6040820160408352808554808352606085019150865f5260209250825f205f5b82811015610f4157815484529284019260019182019101610f25565b50505092019290925292915050565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610f8157634e487b7160e01b5f52601160045260245ffd5b5060010190565b60408152826040820152828460608301375f606084830101525f601f19601f85011682016060838203016020840152610fc46060820185610d21565b9695505050505050565b5f60208284031215610fde575f80fd5b81518015158114610d1a575f80fdfea2646970667358221220045c389b19c7d1c4b7875092b97f6bd0a0169474b8e72a13d55c34deccb34a9764736f6c63430008140033a2646970667358221220e31b81481c11f50718d2a2a0eac5fe6699148616af71f06fe4c7b0cbf1eae3a864736f6c63430008140033",
}

// PollFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PollFactoryMetaData.ABI instead.
var PollFactoryABI = PollFactoryMetaData.ABI

// PollFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PollFactoryMetaData.Bin instead.
var PollFactoryBin = PollFactoryMetaData.Bin

// DeployPollFactory deploys a new Ethereum contract, binding an instance of PollFactory to it.
func DeployPollFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _zkVerifier common.Address, _oracle common.Address) (common.Address, *types.Transaction, *PollFactory, error) {
	parsed, err := PollFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PollFactoryBin), backend, _zkVerifier, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PollFactory{PollFactoryCaller: PollFactoryCaller{contract: contract}, PollFactoryTransactor: PollFactoryTransactor{contract: contract}, PollFactoryFilterer: PollFactoryFilterer{contract: contract}}, nil
}

// PollFactory is an auto generated Go binding around an Ethereum contract.
type PollFactory struct {
	PollFactoryCaller     // Read-only binding to the contract
	PollFactoryTransactor // Write-only binding to the contract
	PollFactoryFilterer   // Log filterer for contract events
}

// PollFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PollFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PollFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PollFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PollFactorySession struct {
	Contract     *PollFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PollFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PollFactoryCallerSession struct {
	Contract *PollFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PollFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PollFactoryTransactorSession struct {
	Contract     *PollFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PollFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PollFactoryRaw struct {
	Contract *PollFactory // Generic contract binding to access the raw methods on
}

// PollFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PollFactoryCallerRaw struct {
	Contract *PollFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PollFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PollFactoryTransactorRaw struct {
	Contract *PollFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPollFactory creates a new instance of PollFactory, bound to a specific deployed contract.
func NewPollFactory(address common.Address, backend bind.ContractBackend) (*PollFactory, error) {
	contract, err := bindPollFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PollFactory{PollFactoryCaller: PollFactoryCaller{contract: contract}, PollFactoryTransactor: PollFactoryTransactor{contract: contract}, PollFactoryFilterer: PollFactoryFilterer{contract: contract}}, nil
}

// NewPollFactoryCaller creates a new read-only instance of PollFactory, bound to a specific deployed contract.
func NewPollFactoryCaller(address common.Address, caller bind.ContractCaller) (*PollFactoryCaller, error) {
	contract, err := bindPollFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PollFactoryCaller{contract: contract}, nil
}

// NewPollFactoryTransactor creates a new write-only instance of PollFactory, bound to a specific deployed contract.
func NewPollFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PollFactoryTransactor, error) {
	contract, err := bindPollFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PollFactoryTransactor{contract: contract}, nil
}

// NewPollFactoryFilterer creates a new log filterer instance of PollFactory, bound to a specific deployed contract.
func NewPollFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PollFactoryFilterer, error) {
	contract, err := bindPollFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PollFactoryFilterer{contract: contract}, nil
}

// bindPollFactory binds a generic wrapper to an already deployed contract.
func bindPollFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PollFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PollFactory *PollFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PollFactory.Contract.PollFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PollFactory *PollFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PollFactory.Contract.PollFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PollFactory *PollFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PollFactory.Contract.PollFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PollFactory *PollFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PollFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PollFactory *PollFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PollFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PollFactory *PollFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PollFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPoll is a free data retrieval call binding the contract method 0x1a8cbcaa.
//
// Solidity: function getPoll(uint256 pollId) view returns(address)
func (_PollFactory *PollFactoryCaller) GetPoll(opts *bind.CallOpts, pollId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "getPoll", pollId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoll is a free data retrieval call binding the contract method 0x1a8cbcaa.
//
// Solidity: function getPoll(uint256 pollId) view returns(address)
func (_PollFactory *PollFactorySession) GetPoll(pollId *big.Int) (common.Address, error) {
	return _PollFactory.Contract.GetPoll(&_PollFactory.CallOpts, pollId)
}

// GetPoll is a free data retrieval call binding the contract method 0x1a8cbcaa.
//
// Solidity: function getPoll(uint256 pollId) view returns(address)
func (_PollFactory *PollFactoryCallerSession) GetPoll(pollId *big.Int) (common.Address, error) {
	return _PollFactory.Contract.GetPoll(&_PollFactory.CallOpts, pollId)
}

// GetPollId is a free data retrieval call binding the contract method 0x46e47f9a.
//
// Solidity: function getPollId(address pollAddress) view returns(uint256)
func (_PollFactory *PollFactoryCaller) GetPollId(opts *bind.CallOpts, pollAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "getPollId", pollAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPollId is a free data retrieval call binding the contract method 0x46e47f9a.
//
// Solidity: function getPollId(address pollAddress) view returns(uint256)
func (_PollFactory *PollFactorySession) GetPollId(pollAddress common.Address) (*big.Int, error) {
	return _PollFactory.Contract.GetPollId(&_PollFactory.CallOpts, pollAddress)
}

// GetPollId is a free data retrieval call binding the contract method 0x46e47f9a.
//
// Solidity: function getPollId(address pollAddress) view returns(uint256)
func (_PollFactory *PollFactoryCallerSession) GetPollId(pollAddress common.Address) (*big.Int, error) {
	return _PollFactory.Contract.GetPollId(&_PollFactory.CallOpts, pollAddress)
}

// GetTotalPolls is a free data retrieval call binding the contract method 0x3ef87d54.
//
// Solidity: function getTotalPolls() view returns(uint256)
func (_PollFactory *PollFactoryCaller) GetTotalPolls(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "getTotalPolls")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPolls is a free data retrieval call binding the contract method 0x3ef87d54.
//
// Solidity: function getTotalPolls() view returns(uint256)
func (_PollFactory *PollFactorySession) GetTotalPolls() (*big.Int, error) {
	return _PollFactory.Contract.GetTotalPolls(&_PollFactory.CallOpts)
}

// GetTotalPolls is a free data retrieval call binding the contract method 0x3ef87d54.
//
// Solidity: function getTotalPolls() view returns(uint256)
func (_PollFactory *PollFactoryCallerSession) GetTotalPolls() (*big.Int, error) {
	return _PollFactory.Contract.GetTotalPolls(&_PollFactory.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PollFactory *PollFactoryCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PollFactory *PollFactorySession) Oracle() (common.Address, error) {
	return _PollFactory.Contract.Oracle(&_PollFactory.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PollFactory *PollFactoryCallerSession) Oracle() (common.Address, error) {
	return _PollFactory.Contract.Oracle(&_PollFactory.CallOpts)
}

// PollCount is a free data retrieval call binding the contract method 0x9207891d.
//
// Solidity: function pollCount() view returns(uint256)
func (_PollFactory *PollFactoryCaller) PollCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "pollCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PollCount is a free data retrieval call binding the contract method 0x9207891d.
//
// Solidity: function pollCount() view returns(uint256)
func (_PollFactory *PollFactorySession) PollCount() (*big.Int, error) {
	return _PollFactory.Contract.PollCount(&_PollFactory.CallOpts)
}

// PollCount is a free data retrieval call binding the contract method 0x9207891d.
//
// Solidity: function pollCount() view returns(uint256)
func (_PollFactory *PollFactoryCallerSession) PollCount() (*big.Int, error) {
	return _PollFactory.Contract.PollCount(&_PollFactory.CallOpts)
}

// PollIds is a free data retrieval call binding the contract method 0xc8cf9ab2.
//
// Solidity: function pollIds(address ) view returns(uint256)
func (_PollFactory *PollFactoryCaller) PollIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "pollIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PollIds is a free data retrieval call binding the contract method 0xc8cf9ab2.
//
// Solidity: function pollIds(address ) view returns(uint256)
func (_PollFactory *PollFactorySession) PollIds(arg0 common.Address) (*big.Int, error) {
	return _PollFactory.Contract.PollIds(&_PollFactory.CallOpts, arg0)
}

// PollIds is a free data retrieval call binding the contract method 0xc8cf9ab2.
//
// Solidity: function pollIds(address ) view returns(uint256)
func (_PollFactory *PollFactoryCallerSession) PollIds(arg0 common.Address) (*big.Int, error) {
	return _PollFactory.Contract.PollIds(&_PollFactory.CallOpts, arg0)
}

// Polls is a free data retrieval call binding the contract method 0xac2f0074.
//
// Solidity: function polls(uint256 ) view returns(address)
func (_PollFactory *PollFactoryCaller) Polls(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "polls", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Polls is a free data retrieval call binding the contract method 0xac2f0074.
//
// Solidity: function polls(uint256 ) view returns(address)
func (_PollFactory *PollFactorySession) Polls(arg0 *big.Int) (common.Address, error) {
	return _PollFactory.Contract.Polls(&_PollFactory.CallOpts, arg0)
}

// Polls is a free data retrieval call binding the contract method 0xac2f0074.
//
// Solidity: function polls(uint256 ) view returns(address)
func (_PollFactory *PollFactoryCallerSession) Polls(arg0 *big.Int) (common.Address, error) {
	return _PollFactory.Contract.Polls(&_PollFactory.CallOpts, arg0)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_PollFactory *PollFactoryCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PollFactory.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_PollFactory *PollFactorySession) ZkVerifier() (common.Address, error) {
	return _PollFactory.Contract.ZkVerifier(&_PollFactory.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_PollFactory *PollFactoryCallerSession) ZkVerifier() (common.Address, error) {
	return _PollFactory.Contract.ZkVerifier(&_PollFactory.CallOpts)
}

// CreatePoll is a paid mutator transaction binding the contract method 0xafea6231.
//
// Solidity: function createPoll(string question, string[] options, uint256 duration, bytes32 voterMerkleRoot) returns(address pollAddress)
func (_PollFactory *PollFactoryTransactor) CreatePoll(opts *bind.TransactOpts, question string, options []string, duration *big.Int, voterMerkleRoot [32]byte) (*types.Transaction, error) {
	return _PollFactory.contract.Transact(opts, "createPoll", question, options, duration, voterMerkleRoot)
}

// CreatePoll is a paid mutator transaction binding the contract method 0xafea6231.
//
// Solidity: function createPoll(string question, string[] options, uint256 duration, bytes32 voterMerkleRoot) returns(address pollAddress)
func (_PollFactory *PollFactorySession) CreatePoll(question string, options []string, duration *big.Int, voterMerkleRoot [32]byte) (*types.Transaction, error) {
	return _PollFactory.Contract.CreatePoll(&_PollFactory.TransactOpts, question, options, duration, voterMerkleRoot)
}

// CreatePoll is a paid mutator transaction binding the contract method 0xafea6231.
//
// Solidity: function createPoll(string question, string[] options, uint256 duration, bytes32 voterMerkleRoot) returns(address pollAddress)
func (_PollFactory *PollFactoryTransactorSession) CreatePoll(question string, options []string, duration *big.Int, voterMerkleRoot [32]byte) (*types.Transaction, error) {
	return _PollFactory.Contract.CreatePoll(&_PollFactory.TransactOpts, question, options, duration, voterMerkleRoot)
}

// PollFactoryPollCreatedIterator is returned from FilterPollCreated and is used to iterate over the raw logs and unpacked data for PollCreated events raised by the PollFactory contract.
type PollFactoryPollCreatedIterator struct {
	Event *PollFactoryPollCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PollFactoryPollCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollFactoryPollCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PollFactoryPollCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PollFactoryPollCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollFactoryPollCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollFactoryPollCreated represents a PollCreated event raised by the PollFactory contract.
type PollFactoryPollCreated struct {
	PollId      *big.Int
	PollAddress common.Address
	Creator     common.Address
	Question    string
	Duration    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPollCreated is a free log retrieval operation binding the contract event 0x5b9a62314cee69a236b75c56d966d19a6903a71666f035bf55a6eec3292ade88.
//
// Solidity: event PollCreated(uint256 indexed pollId, address indexed pollAddress, address indexed creator, string question, uint256 duration)
func (_PollFactory *PollFactoryFilterer) FilterPollCreated(opts *bind.FilterOpts, pollId []*big.Int, pollAddress []common.Address, creator []common.Address) (*PollFactoryPollCreatedIterator, error) {

	var pollIdRule []interface{}
	for _, pollIdItem := range pollId {
		pollIdRule = append(pollIdRule, pollIdItem)
	}
	var pollAddressRule []interface{}
	for _, pollAddressItem := range pollAddress {
		pollAddressRule = append(pollAddressRule, pollAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PollFactory.contract.FilterLogs(opts, "PollCreated", pollIdRule, pollAddressRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &PollFactoryPollCreatedIterator{contract: _PollFactory.contract, event: "PollCreated", logs: logs, sub: sub}, nil
}

// WatchPollCreated is a free log subscription operation binding the contract event 0x5b9a62314cee69a236b75c56d966d19a6903a71666f035bf55a6eec3292ade88.
//
// Solidity: event PollCreated(uint256 indexed pollId, address indexed pollAddress, address indexed creator, string question, uint256 duration)
func (_PollFactory *PollFactoryFilterer) WatchPollCreated(opts *bind.WatchOpts, sink chan<- *PollFactoryPollCreated, pollId []*big.Int, pollAddress []common.Address, creator []common.Address) (event.Subscription, error) {

	var pollIdRule []interface{}
	for _, pollIdItem := range pollId {
		pollIdRule = append(pollIdRule, pollIdItem)
	}
	var pollAddressRule []interface{}
	for _, pollAddressItem := range pollAddress {
		pollAddressRule = append(pollAddressRule, pollAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PollFactory.contract.WatchLogs(opts, "PollCreated", pollIdRule, pollAddressRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollFactoryPollCreated)
				if err := _PollFactory.contract.UnpackLog(event, "PollCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePollCreated is a log parse operation binding the contract event 0x5b9a62314cee69a236b75c56d966d19a6903a71666f035bf55a6eec3292ade88.
//
// Solidity: event PollCreated(uint256 indexed pollId, address indexed pollAddress, address indexed creator, string question, uint256 duration)
func (_PollFactory *PollFactoryFilterer) ParsePollCreated(log types.Log) (*PollFactoryPollCreated, error) {
	event := new(PollFactoryPollCreated)
	if err := _PollFactory.contract.UnpackLog(event, "PollCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
