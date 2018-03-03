/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */
#include "user_mng.hpp"
#include "user_mng.gen.hpp"

#include <eoslib/db.hpp>
#include <eoslib/types.hpp>
#include <eoslib/raw.hpp>


extern "C" {
    void init()  {
   }
   
    void apply( uint64_t code, uint64_t action ) {
        if( code == N(usermng) ) {
            if( action == N(reguser) ) {
                eosio::print("register user begin\n");
                auto req_Info = eosio::current_message<reg_user_req>();                

                //eosio::require_auth( eosio::string_to_name(req_Info.user_name.get_data()) );

                eosio::dump(req_Info);
                bytes b = eosio::raw::pack(req_Info.info);
                uint32_t err = store_str( N(usermng), N(userreginfo), (char *)req_Info.user_name.get_data(), req_Info.user_name.get_size(), (char*)b.data, b.len);

                eosio::print("register user finished\n");           

            }else if( action == N(userlogin) ) {
                eosio::print("user login begin\n");
                auto login_Info = eosio::current_message<user_login>();

                eosio::require_auth( eosio::string_to_name(login_Info.user_name.get_data()) );
                
                eosio::print("user login finished\n");           

            }
            else {
                assert(0, "unknown message");
            }
        }
    }
}
