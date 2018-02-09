/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */
#include <eoslib/eos.hpp>
#include <eoslib/string.hpp>

struct user_basic_Info {
    eosio::string info;
};


/* @abi action reguser
 * @abi table
*/
struct reg_user_req {
    eosio::string user_name;
    user_basic_Info info;
};

