/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */
#include <eoslib/eos.hpp>
#include <eoslib/string.hpp>

/* @abi action login
 * @abi table
*/
struct data_file_info {
    eosio::string user_name;
    eosio::string session_id;
    uint64_t      file_size;
    eosio::string file_name;
    eosio::string file_policy;
    uint64_t      file_number;
    eosio::string signature;
};


struct reg_data_file_req {
    eosio::string  file_hash;
    data_file_info info;
};




